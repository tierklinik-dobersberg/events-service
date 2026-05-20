package dicom

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/dop251/goja"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
)

type Module struct{}

func (*Module) Name() string { return "dicom" }

func (m *Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	exports := vu.Runtime().NewObject()

	tags := createTagObject(vu)
	exports.Set("tag", tags)
	exports.Set("createWorklistDataset", createWorklistDataset)
	exports.Set("write", writeDicomFile)
	exports.Set("read", readDicomFile)

	return exports, nil
}

func init() {
	if err := modules.Register(&Module{}); err != nil {
		panic(fmt.Errorf("failed to register module %w", err))
	}
}

func readDicomFile(path string) (dicom.Dataset, error) {
	f, err := os.Open(path)
	if err != nil {
		return dicom.Dataset{}, fmt.Errorf("failed to open file: %w", err)
	}

	ds, err := dicom.ParseUntilEOF(f, nil, dicom.SkipPixelData())
	if err != nil {
		return dicom.Dataset{}, fmt.Errorf("failed to read DICOM file: %w", err)
	}

	return ds, nil
}

func createTagObject(vu modules.VU) *goja.Object {
	obj := vu.Runtime().NewObject()

	obj.Set("find", findTag)

	for name := range TagNames {
		t, err := findTag(name)
		if err != nil {
			vu.Log().Log(context.Background(), slog.LevelDebug, "failed to find tag", "name", name)
			continue
		}

		obj.Set(name, t)
	}

	return obj
}

type Steps struct {
	StepDescription string `json:"description"`
	StepID          string `json:"id"`
}

type Worklist struct {
	Modality     string  `json:"modality"`
	Steps        []Steps `json:"steps"`
	ScheduledAET string  `json:"scheduled_aet"`

	ClientFirstName string `json:"client_first_name"`
	ClientLastName  string `json:"client_last_name"`
	ClientID        string `json:"client_id"`

	PatientName         string `json:"patient_name"`
	PatientID           string `json:"patient_id"`
	PatientBirthDate    string `json:"patient_birth_date"`
	PatientSex          string `json:"patient_sex"`
	PatientSexNeutered  *bool  `json:"patient_sex_neutered"`
	AdditionalPatientID string `json:"additional_patient_id"`

	RequestingUser string `json:"requesting_user"`
}

func writeDicomFile(ds dicom.Dataset, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	if err := dicom.Write(f, ds, dicom.DefaultMissingTransferSyntax()); err != nil {
		return fmt.Errorf("failed to write dicom dataset: %w", err)
	}

	return nil
}

func createWorklistDataset(cfg Worklist) (dicom.Dataset, error) {
	ds := dicom.Dataset{
		Elements: []*dicom.Element{},
	}

	add := func(t tag.Tag, value any) error {
		var el *dicom.Element
		var err error

		if s, ok := value.(string); ok {
			el, err = dicom.NewElement(t, []string{s})
		} else {
			el, err = dicom.NewElement(t, value)
		}

		if err != nil {
			return err
		}

		ds.Elements = append(ds.Elements, el)

		return nil
	}

	var scheduledProcedureSteps [][]*dicom.Element
	for _, step := range cfg.Steps {
		var scheduledProcedureStep []*dicom.Element
		addStep := func(t tag.Tag, value string) error {
			el, err := dicom.NewElement(t, []string{value})
			if err != nil {
				return err
			}

			scheduledProcedureStep = append(scheduledProcedureStep, el)

			return nil
		}

		if err := addStep(tag.Modality, cfg.Modality); err != nil {
			return ds, fmt.Errorf("failed to configure modality: %w", err)
		}
		if err := addStep(tag.ScheduledStationAETitle, cfg.ScheduledAET); err != nil {
			return ds, fmt.Errorf("failed to configure schedule-aet-title: %w", err)
		}
		if err := addStep(tag.ScheduledProcedureStepDescription, step.StepDescription); err != nil {
			return ds, fmt.Errorf("failed to configure scheduled procedure step description: %w", err)
		}
		if err := addStep(tag.ScheduledProcedureStepID, step.StepID); err != nil {
			return ds, fmt.Errorf("failed to configure scheduled procedure step id: %w", err)
		}

		scheduledProcedureSteps = append(scheduledProcedureSteps, scheduledProcedureStep)
	}

	if err := add(tag.ScheduledProcedureStepSequence, scheduledProcedureSteps); err != nil {
		return ds, fmt.Errorf("failed to configure scheduled procedure sequence id: %w", err)
	}

	name := fmt.Sprintf("%s^%s", cfg.ClientLastName, cfg.ClientFirstName)
	if err := add(tag.ResponsiblePerson, strings.Trim(name, "^")); err != nil {
		return ds, fmt.Errorf("failed to configure responsible person: %w", err)
	}

	if err := add(tag.ResponsiblePersonRole, "Owner"); err != nil {
		return ds, fmt.Errorf("failed to configure responsible person role: %w", err)
	}

	if err := add(tag.RequestingPhysician, cfg.RequestingUser); err != nil {
		return ds, fmt.Errorf("failed to configure requesting physician: %w", err)
	}

	name = cfg.PatientName
	if name == "" {
		name = "Unknown"
	}
	if err := add(tag.PatientName, name); err != nil {
		return ds, fmt.Errorf("failed to configure patient name: %w", err)
	}

	if err := add(tag.PatientID, cfg.PatientID); err != nil {
		return ds, fmt.Errorf("failed to configure patient id: %w", err)
	}

	if err := add(tag.PatientSex, cfg.PatientSex); err != nil {
		return ds, fmt.Errorf("failed to configure patient sex: %w", err)
	}

	if cfg.PatientSexNeutered != nil {
		if *cfg.PatientSexNeutered {
			if err := add(tag.PatientSexNeutered, "ALTERED"); err != nil {
				return ds, fmt.Errorf("failed to configure patient sex neutered: %w", err)
			}
		} else {
			if err := add(tag.PatientSexNeutered, "UNALTERED"); err != nil {
				return ds, fmt.Errorf("failed to configure patient sex neutered: %w", err)
			}
		}
	}

	if err := add(tag.OtherPatientIDs, []string{cfg.AdditionalPatientID, fmt.Sprintf("client:%s", cfg.ClientID)}); err != nil {
		return ds, fmt.Errorf("failed to configure other patient ids: %w", err)
	}

	log.Printf("Creating worklist dataset cfg=%#v ds=%#v", cfg, ds)

	return ds, nil
}

func findTag(name string) (tag.Tag, error) {
	t, err := tag.FindByKeyword(name)
	if err == nil {
		return t.Tag, nil
	}

	t, err = tag.FindByName(name)
	if err != nil {
		return tag.Tag{}, err
	}

	return t.Tag, nil
}
