package dicom

import (
	"encoding/json"
	"fmt"

	"github.com/suyashkumar/dicom"
)

func createDataset() *Dataset {
	return &Dataset{}
}

func fromObject(v map[string]any) (*Dataset, error) {
	ds := createDataset()

	for key, value := range v {
		if err := ds.Add(key, value); err != nil {
			return nil, fmt.Errorf("failed to add element %s: %v", key, err)
		}
	}

	return ds, nil
}

type Dataset struct {
	Elements []*dicom.Element
}

func (ds *Dataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(ds.Elements)
}

func (ds *Dataset) Write(path string) error {
	d := dicom.Dataset{
		Elements: ds.Elements,
	}

	return writeDicomFile(d, path)
}

func (ds *Dataset) Add(name string, value any) error {
	t, err := findTag(name)
	if err != nil {
		return fmt.Errorf("failed to find tag %s: %v", name, err)
	}

	var el *dicom.Element

	switch v := value.(type) {
	case string:
		el, err = dicom.NewElement(t, []string{v})

	case *dicom.Element:
		el = v

	case dicom.Dataset:
		el, err = dicom.NewElement(t, v.Elements)

	case *dicom.Dataset:
		el, err = dicom.NewElement(t, v.Elements)

	case *Dataset:
		el, err = dicom.NewElement(t, v.Elements)

	case Dataset:
		el, err = dicom.NewElement(t, v.Elements)

	default:
		el, err = dicom.NewElement(t, v)
	}

	if err != nil {
		return fmt.Errorf("failed to add element %s|%s: %v", name, t.String(), err)
	}

	ds.Elements = append(ds.Elements, el)

	return nil
}
