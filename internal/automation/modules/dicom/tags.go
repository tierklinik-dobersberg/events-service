package dicom

const (
	FileMetaInfoGroupLength             = "00020000"
	FileMetaInfoVersion                 = "00020001"
	MediaStorageSOPClassUID             = "00020002"
	MediaStorageSOPInstanceUID          = "00020003"
	TransferSyntaxUID                   = "00020010"
	ImplementationClassUID              = "00020012"
	ImplementationVersionName           = "00020013"
	SourceApplicationEntityTitle        = "00020016"
	PrivateInformationCreatorUID        = "00020100"
	PrivateInformation                  = "00020102"
	FileSetID                           = "00041130"
	FileSetDescriptorFileID             = "00041141"
	SpecificCharacterSetOfFile          = "00041142"
	FirstDirectoryRecordOffset          = "00041200"
	LastDirectoryRecordOffset           = "00041202"
	FileSetConsistencyFlag              = "00041212"
	DirectoryRecordSequence             = "00041220"
	OffsetOfNextDirectoryRecord         = "00041400"
	RecordInUseFlag                     = "00041410"
	LowerLevelDirectoryEntityOffset     = "00041420"
	DirectoryRecordType                 = "00041430"
	PrivateRecordUID                    = "00041432"
	ReferencedFileID                    = "00041500"
	MRDRDirectoryRecordOffset           = "00041504"
	ReferencedSOPClassUIDInFile         = "00041510"
	ReferencedSOPInstanceUIDInFile      = "00041511"
	ReferencedTransferSyntaxUIDInFile   = "00041512"
	ReferencedRelatedSOPClassUIDInFile  = "0004151A"
	NumberOfReferences                  = "00041600"
	IdentifyingGroupLength              = "00080000"
	LengthToEnd                         = "00080001"
	SpecificCharacterSet                = "00080005"
	LanguageCodeSequence                = "00080006"
	ImageType                           = "00080008"
	RecognitionCode                     = "00080010"
	InstanceCreationDate                = "00080012"
	InstanceCreationTime                = "00080013"
	InstanceCreatorUID                  = "00080014"
	SOPClassUID                         = "00080016"
	SOPInstanceUID                      = "00080018"
	RelatedGeneralSOPClassUID           = "0008001A"
	OriginalSpecializedSOPClassUID      = "0008001B"
	StudyDate                           = "00080020"
	SeriesDate                          = "00080021"
	AcquisitionDate                     = "00080022"
	ContentDate                         = "00080023"
	OverlayDate                         = "00080024"
	CurveDate                           = "00080025"
	AcquisitionDateTime                 = "0008002A"
	StudyTime                           = "00080030"
	SeriesTime                          = "00080031"
	AcquisitionTime                     = "00080032"
	ContentTime                         = "00080033"
	OverlayTime                         = "00080034"
	CurveTime                           = "00080035"
	DataSetType                         = "00080040"
	DataSetSubtype                      = "00080041"
	NuclearMedicineSeriesType           = "00080042"
	AccessionNumber                     = "00080050"
	QueryRetrieveLevel                  = "00080052"
	RetrieveAETitle                     = "00080054"
	InstanceAvailability                = "00080056"
	FailedSOPInstanceUIDList            = "00080058"
	Modality                            = "00080060"
	ModalitiesInStudy                   = "00080061"
	SOPClassesInStudy                   = "00080062"
	ConversionType                      = "00080064"
	PresentationIntentType              = "00080068"
	Manufacturer                        = "00080070"
	InstitutionName                     = "00080080"
	InstitutionAddress                  = "00080081"
	InstitutionCodeSequence             = "00080082"
	ReferringPhysicianName              = "00080090"
	ReferringPhysicianAddress           = "00080092"
	ReferringPhysicianTelephoneNumber   = "00080094"
	ReferringPhysicianIDSequence        = "00080096"
	CodeValue                           = "00080100"
	CodingSchemeDesignator              = "00080102"
	CodingSchemeVersion                 = "00080103"
	CodeMeaning                         = "00080104"
	MappingResource                     = "00080105"
	ContextGroupVersion                 = "00080106"
	ContextGroupLocalVersion            = "00080107"
	ContextGroupExtensionFlag           = "0008010B"
	CodingSchemeUID                     = "0008010C"
	ContextGroupExtensionCreatorUID     = "0008010D"
	ContextIdentifier                   = "0008010F"
	CodingSchemeIDSequence              = "00080110"
	CodingSchemeRegistry                = "00080112"
	CodingSchemeExternalID              = "00080114"
	CodingSchemeName                    = "00080115"
	CodingSchemeResponsibleOrganization = "00080116"
	ContextUID                          = "00080117"
	TimezoneOffsetFromUTC               = "00080201"
	NetworkID                           = "00081000"
	StationName                         = "00081010"
	StudyDescription                    = "00081030"
	ProcedureCodeSequence               = "00081032"
	SeriesDescription                   = "0008103E"
	InstitutionalDepartmentName         = "00081040"
	PhysiciansOfRecord                  = "00081048"
	PhysiciansOfRecordIDSequence        = "00081049"
	PerformingPhysicianName             = "00081050"
	PerformingPhysicianIDSequence       = "00081052"
	NameOfPhysicianReadingStudy         = "00081060"
	PhysicianReadingStudyIDSequence     = "00081062"
	OperatorsName                       = "00081070"
	OperatorIDSequence                  = "00081072"
	AdmittingDiagnosesDescription       = "00081080"
	AdmittingDiagnosesCodeSequence      = "00081084"
	ManufacturersModelName              = "00081090"
	ReferencedResultsSequence           = "00081100"
	ReferencedStudySequence             = "00081110"
	ReferencedProcedureStepSequence     = "00081111"
	ReferencedSeriesSequence            = "00081115"
	ReferencedPatientSequence           = "00081120"
	ReferencedVisitSequence             = "00081125"
	ReferencedOverlaySequence           = "00081130"
	ReferencedWaveformSequence          = "0008113A"
	ReferencedImageSequence             = "00081140"
	ReferencedCurveSequence             = "00081145"
	ReferencedInstanceSequence          = "0008114A"
	ReferencedSOPClassUID               = "00081150"
	ReferencedSOPInstanceUID            = "00081155"
	SOPClassesSupported                 = "0008115A"
	ReferencedFrameNumber               = "00081160"
	SimpleFrameList                     = "00081161"
	CalculatedFrameList                 = "00081162"
	TimeRange                           = "00081163"
	FrameExtractionSequence             = "00081164"
	RetrieveURL                         = "00081190"
	TransactionUID                      = "00081195"
	FailureReason                       = "00081197"
	FailedSOPSequence                   = "00081198"
	ReferencedSOPSequence               = "00081199"
	OtherReferencedStudiesSequence      = "00081200"
	RelatedSeriesSequence               = "00081250"
	LossyImageCompression               = "00082110"
	DerivationDescription               = "00082111"
	SourceImageSequence                 = "00082112"
	StageName                           = "00082120"
	StageNumber                         = "00082122"
	NumberOfStages                      = "00082124"
	ViewName                            = "00082127"
	ViewNumber                          = "00082128"
	NumberOfEventTimers                 = "00082129"
	NumberOfViewsInStage                = "0008212A"
	EventElapsedTimes                   = "00082130"
	EventTimerNames                     = "00082132"
	EventTimerSequence                  = "00082133"
	EventTimeOffset                     = "00082134"
	EventCodeSequence                   = "00082135"
	StartTrim                           = "00082142"
	StopTrim                            = "00082143"
	RecommendedDisplayFrameRate         = "00082144"
	TransducerPosition                  = "00082200"
	TransducerOrientation               = "00082204"
	AnatomicStructure                   = "00082208"
	AnatomicRegionSequence              = "00082218"
	AnatomicRegionModifierSequence      = "00082220"
	PrimaryAnatomicStructureSequence    = "00082228"
	AnatomicStructureOrRegionSequence   = "00082229"
	AnatomicStructureModifierSequence   = "00082230"
	TransducerPositionSequence          = "00082240"
	TransducerPositionModifierSequence  = "00082242"
	TransducerOrientationSequence       = "00082244"
	TransducerOrientationModifierSeq    = "00082246"
	AnatomicEntrancePortalCodeSeqTrial  = "00082253"
	AnatomicApproachDirCodeSeqTrial     = "00082255"
	AnatomicPerspectiveDescrTrial       = "00082256"
	AnatomicPerspectiveCodeSeqTrial     = "00082257"
	AlternateRepresentationSequence     = "00083001"
	IrradiationEventUID                 = "00083010"
	IdentifyingComments                 = "00084000"
	FrameType                           = "00089007"
	ReferencedImageEvidenceSequence     = "00089092"
	ReferencedRawDataSequence           = "00089121"
	CreatorVersionUID                   = "00089123"
	DerivationImageSequence             = "00089124"
	SourceImageEvidenceSequence         = "00089154"
	PixelPresentation                   = "00089205"
	VolumetricProperties                = "00089206"
	VolumeBasedCalculationTechnique     = "00089207"
	ComplexImageComponent               = "00089208"
	AcquisitionContrast                 = "00089209"
	DerivationCodeSequence              = "00089215"
	GrayscalePresentationStateSequence  = "00089237"
	ReferencedOtherPlaneSequence        = "00089410"
	FrameDisplaySequence                = "00089458"
	RecommendedDisplayFrameRateInFloat  = "00089459"
	SkipFrameRangeFlag                  = "00089460"
	FullFidelity                        = "00091001"
	SuiteID                             = "00091002"
	ProductID                           = "00091004"
	ImageActualDate                     = "00091027"
	ServiceID                           = "00091030"
	MobileLocationNumber                = "00091031"
	EquipmentUID                        = "000910E3"
	GenesisVersionNow                   = "000910E6"
	ExamRecordChecksum                  = "000910E7"
	ActualSeriesDataTimeStamp           = "000910E9"
	PatientGroupLength                  = "00100000"
	PatientName                         = "00100010"
	PatientID                           = "00100020"
	IssuerOfPatientID                   = "00100021"
	TypeOfPatientID                     = "00100022"
	PatientBirthDate                    = "00100030"
	PatientBirthTime                    = "00100032"
	PatientSex                          = "00100040"
	PatientInsurancePlanCodeSequence    = "00100050"
	PatientPrimaryLanguageCodeSeq       = "00100101"
	PatientPrimaryLanguageCodeModSeq    = "00100102"
	OtherPatientIDs                     = "00101000"
	OtherPatientNames                   = "00101001"
	OtherPatientIDsSequence             = "00101002"
	PatientBirthName                    = "00101005"
	PatientAge                          = "00101010"
	PatientSize                         = "00101020"
	PatientWeight                       = "00101030"
	PatientAddress                      = "00101040"
	InsurancePlanIdentification         = "00101050"
	PatientMotherBirthName              = "00101060"
	MilitaryRank                        = "00101080"
	BranchOfService                     = "00101081"
	MedicalRecordLocator                = "00101090"
	MedicalAlerts                       = "00102000"
	Allergies                           = "00102110"
	CountryOfResidence                  = "00102150"
	RegionOfResidence                   = "00102152"
	PatientTelephoneNumbers             = "00102154"
	EthnicGroup                         = "00102160"
	Occupation                          = "00102180"
	SmokingStatus                       = "001021A0"
	AdditionalPatientHistory            = "001021B0"
	PregnancyStatus                     = "001021C0"
	LastMenstrualDate                   = "001021D0"
	PatientReligiousPreference          = "001021F0"
	PatientSpeciesDescription           = "00102201"
	PatientSpeciesCodeSequence          = "00102202"
	PatientSexNeutered                  = "00102203"
	AnatomicalOrientationType           = "00102210"
	PatientBreedDescription             = "00102292"
	PatientBreedCodeSequence            = "00102293"
	BreedRegistrationSequence           = "00102294"
	BreedRegistrationNumber             = "00102295"
	BreedRegistryCodeSequence           = "00102296"
	ResponsiblePerson                   = "00102297"
	ResponsiblePersonRole               = "00102298"
	ResponsibleOrganization             = "00102299"
	PatientComments                     = "00104000"
	ExaminedBodyThickness               = "00109431"
	PatientStatus                       = "00111010"
	ClinicalTrialSponsorName            = "00120010"
	ClinicalTrialProtocolID             = "00120020"
	ClinicalTrialProtocolName           = "00120021"
	ClinicalTrialSiteID                 = "00120030"
	ClinicalTrialSiteName               = "00120031"
	ClinicalTrialSubjectID              = "00120040"
	ClinicalTrialSubjectReadingID       = "00120042"
	ClinicalTrialTimePointID            = "00120050"
	ClinicalTrialTimePointDescription   = "00120051"
	ClinicalTrialCoordinatingCenter     = "00120060"
	PatientIdentityRemoved              = "00120062"
	DeidentificationMethod              = "00120063"
	DeidentificationMethodCodeSequence  = "00120064"
	ClinicalTrialSeriesID               = "00120071"
	ClinicalTrialSeriesDescription      = "00120072"
	DistributionType                    = "00120084"
	ConsentForDistributionFlag          = "00120085"
	AcquisitionGroupLength              = "00180000"
	ContrastBolusAgent                  = "00180010"
	ContrastBolusAgentSequence          = "00180012"
	ContrastBolusAdministrationRoute    = "00180014"
	BodyPartExamined                    = "00180015"
	ScanningSequence                    = "00180020"
	SequenceVariant                     = "00180021"
	ScanOptions                         = "00180022"
	MRAcquisitionType                   = "00180023"
	SequenceName                        = "00180024"
	AngioFlag                           = "00180025"
	InterventionDrugInformationSeq      = "00180026"
	InterventionDrugStopTime            = "00180027"
	InterventionDrugDose                = "00180028"
	InterventionDrugSequence            = "00180029"
	AdditionalDrugSequence              = "0018002A"
	Radionuclide                        = "00180030"
	Radiopharmaceutical                 = "00180031"
	EnergyWindowCenterline              = "00180032"
	EnergyWindowTotalWidth              = "00180033"
	InterventionDrugName                = "00180034"
	InterventionDrugStartTime           = "00180035"
	InterventionSequence                = "00180036"
	TherapyType                         = "00180037"
	InterventionStatus                  = "00180038"
	TherapyDescription                  = "00180039"
	InterventionDescription             = "0018003A"
	CineRate                            = "00180040"
	InitialCineRunState                 = "00180042"
	SliceThickness                      = "00180050"
	KVP                                 = "00180060"
	CountsAccumulated                   = "00180070"
	AcquisitionTerminationCondition     = "00180071"
	EffectiveDuration                   = "00180072"
	AcquisitionStartCondition           = "00180073"
	AcquisitionStartConditionData       = "00180074"
	AcquisitionEndConditionData         = "00180075"
	RepetitionTime                      = "00180080"
	EchoTime                            = "00180081"
	InversionTime                       = "00180082"
	NumberOfAverages                    = "00180083"
	ImagingFrequency                    = "00180084"
	ImagedNucleus                       = "00180085"
	EchoNumber                          = "00180086"
	MagneticFieldStrength               = "00180087"
	SpacingBetweenSlices                = "00180088"
	NumberOfPhaseEncodingSteps          = "00180089"
	DataCollectionDiameter              = "00180090"
	EchoTrainLength                     = "00180091"
	PercentSampling                     = "00180093"
	PercentPhaseFieldOfView             = "00180094"
	PixelBandwidth                      = "00180095"
	DeviceSerialNumber                  = "00181000"
	DeviceUID                           = "00181002"
	DeviceID                            = "00181003"
	PlateID                             = "00181004"
	GeneratorID                         = "00181005"
	GridID                              = "00181006"
	CassetteID                          = "00181007"
	GantryID                            = "00181008"
	SecondaryCaptureDeviceID            = "00181010"
	HardcopyCreationDeviceID            = "00181011"
	DateOfSecondaryCapture              = "00181012"
	TimeOfSecondaryCapture              = "00181014"
	SecondaryCaptureDeviceManufacturer  = "00181016"
	HardcopyDeviceManufacturer          = "00181017"
	SecondaryCaptureDeviceModelName     = "00181018"
	SecondaryCaptureDeviceSoftwareVers  = "00181019"
	HardcopyDeviceSoftwareVersion       = "0018101A"
	HardcopyDeviceModelName             = "0018101B"
	SoftwareVersion                     = "00181020"
	VideoImageFormatAcquired            = "00181022"
	DigitalImageFormatAcquired          = "00181023"
	ProtocolName                        = "00181030"
	ContrastBolusRoute                  = "00181040"
	ContrastBolusVolume                 = "00181041"
	ContrastBolusStartTime              = "00181042"
	ContrastBolusStopTime               = "00181043"
	ContrastBolusTotalDose              = "00181044"
	SyringeCounts                       = "00181045"
	ContrastFlowRate                    = "00181046"
	ContrastFlowDuration                = "00181047"
	ContrastBolusIngredient             = "00181048"
	ContrastBolusConcentration          = "00181049"
	SpatialResolution                   = "00181050"
	TriggerTime                         = "00181060"
	TriggerSourceOrType                 = "00181061"
	NominalInterval                     = "00181062"
	FrameTime                           = "00181063"
	CardiacFramingType                  = "00181064"
	FrameTimeVector                     = "00181065"
	FrameDelay                          = "00181066"
	ImageTriggerDelay                   = "00181067"
	MultiplexGroupTimeOffset            = "00181068"
	TriggerTimeOffset                   = "00181069"
	SynchronizationTrigger              = "0018106A"
	SynchronizationChannel              = "0018106C"
	TriggerSamplePosition               = "0018106E"
	RadiopharmaceuticalRoute            = "00181070"
	RadiopharmaceuticalVolume           = "00181071"
	RadiopharmaceuticalStartTime        = "00181072"
	RadiopharmaceuticalStopTime         = "00181073"
	RadionuclideTotalDose               = "00181074"
	RadionuclideHalfLife                = "00181075"
	RadionuclidePositronFraction        = "00181076"
	RadiopharmaceuticalSpecActivity     = "00181077"
	RadiopharmaceuticalStartDateTime    = "00181078"
	RadiopharmaceuticalStopDateTime     = "00181079"
	BeatRejectionFlag                   = "00181080"
	LowRRValue                          = "00181081"
	HighRRValue                         = "00181082"
	IntervalsAcquired                   = "00181083"
	IntervalsRejected                   = "00181084"
	PVCRejection                        = "00181085"
	SkipBeats                           = "00181086"
	HeartRate                           = "00181088"
	CardiacNumberOfImages               = "00181090"
	TriggerWindow                       = "00181094"
	ReconstructionDiameter              = "00181100"
	DistanceSourceToDetector            = "00181110"
	DistanceSourceToPatient             = "00181111"
	EstimatedRadiographicMagnification  = "00181114"
	GantryDetectorTilt                  = "00181120"
	GantryDetectorSlew                  = "00181121"
	TableHeight                         = "00181130"
	TableTraverse                       = "00181131"
	TableMotion                         = "00181134"
	TableVerticalIncrement              = "00181135"
	TableLateralIncrement               = "00181136"
	TableLongitudinalIncrement          = "00181137"
	TableAngle                          = "00181138"
	TableType                           = "0018113A"
	RotationDirection                   = "00181140"
	AngularPosition                     = "00181141"
	RadialPosition                      = "00181142"
	ScanArc                             = "00181143"
	AngularStep                         = "00181144"
	CenterOfRotationOffset              = "00181145"
	RotationOffset                      = "00181146"
	FieldOfViewShape                    = "00181147"
	FieldOfViewDimensions               = "00181149"
	ExposureTime                        = "00181150"
	XRayTubeCurrent                     = "00181151"
	Exposure                            = "00181152"
	ExposureInMicroAmpSec               = "00181153"
	AveragePulseWidth                   = "00181154"
	RadiationSetting                    = "00181155"
	RectificationType                   = "00181156"
	RadiationMode                       = "0018115A"
	ImageAreaDoseProduct                = "0018115E"
	FilterType                          = "00181160"
	TypeOfFilters                       = "00181161"
	IntensifierSize                     = "00181162"
	ImagerPixelSpacing                  = "00181164"
	Grid                                = "00181166"
	GeneratorPower                      = "00181170"
	CollimatorGridName                  = "00181180"
	CollimatorType                      = "00181181"
	FocalDistance                       = "00181182"
	XFocusCenter                        = "00181183"
	YFocusCenter                        = "00181184"
	FocalSpots                          = "00181190"
	AnodeTargetMaterial                 = "00181191"
	BodyPartThickness                   = "001811A0"
	CompressionForce                    = "001811A2"
	DateOfLastCalibration               = "00181200"
	TimeOfLastCalibration               = "00181201"
	ConvolutionKernel                   = "00181210"
	UpperLowerPixelValues               = "00181240"
	ActualFrameDuration                 = "00181242"
	CountRate                           = "00181243"
	PreferredPlaybackSequencing         = "00181244"
	ReceiveCoilName                     = "00181250"
	TransmitCoilName                    = "00181251"
	PlateType                           = "00181260"
	PhosphorType                        = "00181261"
	ScanVelocity                        = "00181300"
	WholeBodyTechnique                  = "00181301"
	ScanLength                          = "00181302"
	AcquisitionMatrix                   = "00181310"
	InPlanePhaseEncodingDirection       = "00181312"
	FlipAngle                           = "00181314"
	VariableFlipAngleFlag               = "00181315"
	SAR                                 = "00181316"
	DBDt                                = "00181318"
	AcquisitionDeviceProcessingDescr    = "00181400"
	AcquisitionDeviceProcessingCode     = "00181401"
	CassetteOrientation                 = "00181402"
	CassetteSize                        = "00181403"
	ExposuresOnPlate                    = "00181404"
	RelativeXRayExposure                = "00181405"
	ColumnAngulation                    = "00181450"
	TomoLayerHeight                     = "00181460"
	TomoAngle                           = "00181470"
	TomoTime                            = "00181480"
	TomoType                            = "00181490"
	TomoClass                           = "00181491"
	NumberOfTomosynthesisSourceImages   = "00181495"
	PositionerMotion                    = "00181500"
	PositionerType                      = "00181508"
	PositionerPrimaryAngle              = "00181510"
	PositionerSecondaryAngle            = "00181511"
	PositionerPrimaryAngleIncrement     = "00181520"
	PositionerSecondaryAngleIncrement   = "00181521"
	DetectorPrimaryAngle                = "00181530"
	DetectorSecondaryAngle              = "00181531"
	ShutterShape                        = "00181600"
	ShutterLeftVerticalEdge             = "00181602"
	ShutterRightVerticalEdge            = "00181604"
	ShutterUpperHorizontalEdge          = "00181606"
	ShutterLowerHorizontalEdge          = "00181608"
	CenterOfCircularShutter             = "00181610"
	RadiusOfCircularShutter             = "00181612"
	VerticesOfPolygonalShutter          = "00181620"
	ShutterPresentationValue            = "00181622"
	ShutterOverlayGroup                 = "00181623"
	ShutterPresentationColorCIELabVal   = "00181624"
	CollimatorShape                     = "00181700"
	CollimatorLeftVerticalEdge          = "00181702"
	CollimatorRightVerticalEdge         = "00181704"
	CollimatorUpperHorizontalEdge       = "00181706"
	CollimatorLowerHorizontalEdge       = "00181708"
	CenterOfCircularCollimator          = "00181710"
	RadiusOfCircularCollimator          = "00181712"
	VerticesOfPolygonalCollimator       = "00181720"
	AcquisitionTimeSynchronized         = "00181800"
	TimeSource                          = "00181801"
	TimeDistributionProtocol            = "00181802"
	NTPSourceAddress                    = "00181803"
	PageNumberVector                    = "00182001"
	FrameLabelVector                    = "00182002"
	FramePrimaryAngleVector             = "00182003"
	FrameSecondaryAngleVector           = "00182004"
	SliceLocationVector                 = "00182005"
	DisplayWindowLabelVector            = "00182006"
	NominalScannedPixelSpacing          = "00182010"
	DigitizingDeviceTransportDirection  = "00182020"
	RotationOfScannedFilm               = "00182030"
	IVUSAcquisition                     = "00183100"
	IVUSPullbackRate                    = "00183101"
	IVUSGatedRate                       = "00183102"
	IVUSPullbackStartFrameNumber        = "00183103"
	IVUSPullbackStopFrameNumber         = "00183104"
	LesionNumber                        = "00183105"
	AcquisitionComments                 = "00184000"
	OutputPower                         = "00185000"
	TransducerData                      = "00185010"
	FocusDepth                          = "00185012"
	ProcessingFunction                  = "00185020"
	PostprocessingFunction              = "00185021"
	MechanicalIndex                     = "00185022"
	BoneThermalIndex                    = "00185024"
	CranialThermalIndex                 = "00185026"
	SoftTissueThermalIndex              = "00185027"
	SoftTissueFocusThermalIndex         = "00185028"
	SoftTissueSurfaceThermalIndex       = "00185029"
	DynamicRange                        = "00185030"
	TotalGain                           = "00185040"
	DepthOfScanField                    = "00185050"
	PatientPosition                     = "00185100"
	ViewPosition                        = "00185101"
	ProjectionEponymousNameCodeSeq      = "00185104"
	ImageTransformationMatrix           = "00185210"
	ImageTranslationVector              = "00185212"
	Sensitivity                         = "00186000"
	SequenceOfUltrasoundRegions         = "00186011"
	RegionSpatialFormat                 = "00186012"
	RegionDataType                      = "00186014"
	RegionFlags                         = "00186016"
	RegionLocationMinX0                 = "00186018"
	RegionLocationMinY0                 = "0018601A"
	RegionLocationMaxX1                 = "0018601C"
	RegionLocationMaxY1                 = "0018601E"
	ReferencePixelX0                    = "00186020"
	ReferencePixelY0                    = "00186022"
	PhysicalUnitsXDirection             = "00186024"
	PhysicalUnitsYDirection             = "00186026"
	ReferencePixelPhysicalValueX        = "00186028"
	ReferencePixelPhysicalValueY        = "0018602A"
	PhysicalDeltaX                      = "0018602C"
	PhysicalDeltaY                      = "0018602E"
	TransducerFrequency                 = "00186030"
	TransducerType                      = "00186031"
	PulseRepetitionFrequency            = "00186032"
	DopplerCorrectionAngle              = "00186034"
	SteeringAngle                       = "00186036"
	DopplerSampleVolumeXPosRetired      = "00186038"
	DopplerSampleVolumeXPosition        = "00186039"
	DopplerSampleVolumeYPosRetired      = "0018603A"
	DopplerSampleVolumeYPosition        = "0018603B"
	TMLinePositionX0Retired             = "0018603C"
	TMLinePositionX0                    = "0018603D"
	TMLinePositionY0Retired             = "0018603E"
	TMLinePositionY0                    = "0018603F"
	TMLinePositionX1Retired             = "00186040"
	TMLinePositionX1                    = "00186041"
	TMLinePositionY1Retired             = "00186042"
	TMLinePositionY1                    = "00186043"
	PixelComponentOrganization          = "00186044"
	PixelComponentMask                  = "00186046"
	PixelComponentRangeStart            = "00186048"
	PixelComponentRangeStop             = "0018604A"
	PixelComponentPhysicalUnits         = "0018604C"
	PixelComponentDataType              = "0018604E"
	NumberOfTableBreakPoints            = "00186050"
	TableOfXBreakPoints                 = "00186052"
	TableOfYBreakPoints                 = "00186054"
	NumberOfTableEntries                = "00186056"
	TableOfPixelValues                  = "00186058"
	TableOfParameterValues              = "0018605A"
	RWaveTimeVector                     = "00186060"
	DetectorConditionsNominalFlag       = "00187000"
	DetectorTemperature                 = "00187001"
	DetectorType                        = "00187004"
	DetectorConfiguration               = "00187005"
	DetectorDescription                 = "00187006"
	DetectorMode                        = "00187008"
	DetectorID                          = "0018700A"
	DateOfLastDetectorCalibration       = "0018700C"
	TimeOfLastDetectorCalibration       = "0018700E"
	DetectorExposuresSinceCalibration   = "00187010"
	DetectorExposuresSinceManufactured  = "00187011"
	DetectorTimeSinceLastExposure       = "00187012"
	DetectorActiveTime                  = "00187014"
	DetectorActiveOffsetFromExposure    = "00187016"
	DetectorBinning                     = "0018701A"
	DetectorElementPhysicalSize         = "00187020"
	DetectorElementSpacing              = "00187022"
	DetectorActiveShape                 = "00187024"
	DetectorActiveDimensions            = "00187026"
	DetectorActiveOrigin                = "00187028"
	DetectorManufacturerName            = "0018702A"
	DetectorManufacturersModelName      = "0018702B"
	FieldOfViewOrigin                   = "00187030"
	FieldOfViewRotation                 = "00187032"
	FieldOfViewHorizontalFlip           = "00187034"
	GridAbsorbingMaterial               = "00187040"
	GridSpacingMaterial                 = "00187041"
	GridThickness                       = "00187042"
	GridPitch                           = "00187044"
	GridAspectRatio                     = "00187046"
	GridPeriod                          = "00187048"
	GridFocalDistance                   = "0018704C"
	FilterMaterial                      = "00187050"
	FilterThicknessMinimum              = "00187052"
	FilterThicknessMaximum              = "00187054"
	ExposureControlMode                 = "00187060"
	ExposureControlModeDescription      = "00187062"
	ExposureStatus                      = "00187064"
	PhototimerSetting                   = "00187065"
	ExposureTimeInMicroSec              = "00188150"
	XRayTubeCurrentInMicroAmps          = "00188151"
	ContentQualification                = "00189004"
	PulseSequenceName                   = "00189005"
	MRImagingModifierSequence           = "00189006"
	EchoPulseSequence                   = "00189008"
	InversionRecovery                   = "00189009"
	FlowCompensation                    = "00189010"
	MultipleSpinEcho                    = "00189011"
	MultiPlanarExcitation               = "00189012"
	PhaseContrast                       = "00189014"
	TimeOfFlightContrast                = "00189015"
	Spoiling                            = "00189016"
	SteadyStatePulseSequence            = "00189017"
	EchoPlanarPulseSequence             = "00189018"
	TagAngleFirstAxis                   = "00189019"
	MagnetizationTransfer               = "00189020"
	T2Preparation                       = "00189021"
	BloodSignalNulling                  = "00189022"
	SaturationRecovery                  = "00189024"
	SpectrallySelectedSuppression       = "00189025"
	SpectrallySelectedExcitation        = "00189026"
	SpatialPresaturation                = "00189027"
	Tagging                             = "00189028"
	OversamplingPhase                   = "00189029"
	TagSpacingFirstDimension            = "00189030"
	GeometryOfKSpaceTraversal           = "00189032"
	SegmentedKSpaceTraversal            = "00189033"
	RectilinearPhaseEncodeReordering    = "00189034"
	TagThickness                        = "00189035"
	PartialFourierDirection             = "00189036"
	CardiacSynchronizationTechnique     = "00189037"
	ReceiveCoilManufacturerName         = "00189041"
	MRReceiveCoilSequence               = "00189042"
	ReceiveCoilType                     = "00189043"
	QuadratureReceiveCoil               = "00189044"
	MultiCoilDefinitionSequence         = "00189045"
	MultiCoilConfiguration              = "00189046"
	MultiCoilElementName                = "00189047"
	MultiCoilElementUsed                = "00189048"
	MRTransmitCoilSequence              = "00189049"
	TransmitCoilManufacturerName        = "00189050"
	TransmitCoilType                    = "00189051"
	SpectralWidth                       = "00189052"
	ChemicalShiftReference              = "00189053"
	VolumeLocalizationTechnique         = "00189054"
	MRAcquisitionFrequencyEncodeSteps   = "00189058"
	Decoupling                          = "00189059"
	DecoupledNucleus                    = "00189060"
	DecouplingFrequency                 = "00189061"
	DecouplingMethod                    = "00189062"
	DecouplingChemicalShiftReference    = "00189063"
	KSpaceFiltering                     = "00189064"
	TimeDomainFiltering                 = "00189065"
	NumberOfZeroFills                   = "00189066"
	BaselineCorrection                  = "00189067"
	ParallelReductionFactorInPlane      = "00189069"
	CardiacRRIntervalSpecified          = "00189070"
	AcquisitionDuration                 = "00189073"
	FrameAcquisitionDateTime            = "00189074"
	DiffusionDirectionality             = "00189075"
	DiffusionGradientDirectionSequence  = "00189076"
	ParallelAcquisition                 = "00189077"
	ParallelAcquisitionTechnique        = "00189078"
	InversionTimes                      = "00189079"
	MetaboliteMapDescription            = "00189080"
	PartialFourier                      = "00189081"
	EffectiveEchoTime                   = "00189082"
	MetaboliteMapCodeSequence           = "00189083"
	ChemicalShiftSequence               = "00189084"
	CardiacSignalSource                 = "00189085"
	DiffusionBValue                     = "00189087"
	DiffusionGradientOrientation        = "00189089"
	VelocityEncodingDirection           = "00189090"
	VelocityEncodingMinimumValue        = "00189091"
	NumberOfKSpaceTrajectories          = "00189093"
	CoverageOfKSpace                    = "00189094"
	SpectroscopyAcquisitionPhaseRows    = "00189095"
	ParallelReductFactorInPlaneRetired  = "00189096"
	TransmitterFrequency                = "00189098"
	ResonantNucleus                     = "00189100"
	FrequencyCorrection                 = "00189101"
	MRSpectroscopyFOVGeometrySequence   = "00189103"
	SlabThickness                       = "00189104"
	SlabOrientation                     = "00189105"
	MidSlabPosition                     = "00189106"
	MRSpatialSaturationSequence         = "00189107"
	MRTimingAndRelatedParametersSeq     = "00189112"
	MREchoSequence                      = "00189114"
	MRModifierSequence                  = "00189115"
	MRDiffusionSequence                 = "00189117"
	CardiacTriggerSequence              = "00189118"
	MRAveragesSequence                  = "00189119"
	MRFOVGeometrySequence               = "00189125"
	VolumeLocalizationSequence          = "00189126"
	SpectroscopyAcquisitionDataColumns  = "00189127"
	DiffusionAnisotropyType             = "00189147"
	FrameReferenceDateTime              = "00189151"
	MRMetaboliteMapSequence             = "00189152"
	ParallelReductionFactorOutOfPlane   = "00189155"
	SpectroscopyOutOfPlanePhaseSteps    = "00189159"
	BulkMotionStatus                    = "00189166"
	ParallelReductionFactSecondInPlane  = "00189168"
	CardiacBeatRejectionTechnique       = "00189169"
	RespiratoryMotionCompTechnique      = "00189170"
	RespiratorySignalSource             = "00189171"
	BulkMotionCompensationTechnique     = "00189172"
	BulkMotionSignalSource              = "00189173"
	ApplicableSafetyStandardAgency      = "00189174"
	ApplicableSafetyStandardDescr       = "00189175"
	OperatingModeSequence               = "00189176"
	OperatingModeType                   = "00189177"
	OperatingMode                       = "00189178"
	SpecificAbsorptionRateDefinition    = "00189179"
	GradientOutputType                  = "00189180"
	SpecificAbsorptionRateValue         = "00189181"
	GradientOutput                      = "00189182"
	FlowCompensationDirection           = "00189183"
	TaggingDelay                        = "00189184"
	RespiratoryMotionCompTechDescr      = "00189185"
	RespiratorySignalSourceID           = "00189186"
	ChemicalShiftsMinIntegrateLimitHz   = "00189195"
	ChemicalShiftsMaxIntegrateLimitHz   = "00189196"
	MRVelocityEncodingSequence          = "00189197"
	FirstOrderPhaseCorrection           = "00189198"
	WaterReferencedPhaseCorrection      = "00189199"
	MRSpectroscopyAcquisitionType       = "00189200"
	RespiratoryCyclePosition            = "00189214"
	VelocityEncodingMaximumValue        = "00189217"
	TagSpacingSecondDimension           = "00189218"
	TagAngleSecondAxis                  = "00189219"
	FrameAcquisitionDuration            = "00189220"
	MRImageFrameTypeSequence            = "00189226"
	MRSpectroscopyFrameTypeSequence     = "00189227"
	MRAcqPhaseEncodingStepsInPlane      = "00189231"
	MRAcqPhaseEncodingStepsOutOfPlane   = "00189232"
	SpectroscopyAcqPhaseColumns         = "00189234"
	CardiacCyclePosition                = "00189236"
	SpecificAbsorptionRateSequence      = "00189239"
	RFEchoTrainLength                   = "00189240"
	GradientEchoTrainLength             = "00189241"
	ChemicalShiftsMinIntegrateLimitPPM  = "00189295"
	ChemicalShiftsMaxIntegrateLimitPPM  = "00189296"
	CTAcquisitionTypeSequence           = "00189301"
	AcquisitionType                     = "00189302"
	TubeAngle                           = "00189303"
	CTAcquisitionDetailsSequence        = "00189304"
	RevolutionTime                      = "00189305"
	SingleCollimationWidth              = "00189306"
	TotalCollimationWidth               = "00189307"
	CTTableDynamicsSequence             = "00189308"
	TableSpeed                          = "00189309"
	TableFeedPerRotation                = "00189310"
	SpiralPitchFactor                   = "00189311"
	CTGeometrySequence                  = "00189312"
	DataCollectionCenterPatient         = "00189313"
	CTReconstructionSequence            = "00189314"
	ReconstructionAlgorithm             = "00189315"
	ConvolutionKernelGroup              = "00189316"
	ReconstructionFieldOfView           = "00189317"
	ReconstructionTargetCenterPatient   = "00189318"
	ReconstructionAngle                 = "00189319"
	ImageFilter                         = "00189320"
	CTExposureSequence                  = "00189321"
	ReconstructionPixelSpacing          = "00189322"
	ExposureModulationType              = "00189323"
	EstimatedDoseSaving                 = "00189324"
	CTXRayDetailsSequence               = "00189325"
	CTPositionSequence                  = "00189326"
	TablePosition                       = "00189327"
	ExposureTimeInMilliSec              = "00189328"
	CTImageFrameTypeSequence            = "00189329"
	XRayTubeCurrentInMilliAmps          = "00189330"
	ExposureInMilliAmpSec               = "00189332"
	ConstantVolumeFlag                  = "00189333"
	FluoroscopyFlag                     = "00189334"
	SourceToDataCollectionCenterDist    = "00189335"
	ContrastBolusAgentNumber            = "00189337"
	ContrastBolusIngredientCodeSeq      = "00189338"
	ContrastAdministrationProfileSeq    = "00189340"
	ContrastBolusUsageSequence          = "00189341"
	ContrastBolusAgentAdministered      = "00189342"
	ContrastBolusAgentDetected          = "00189343"
	ContrastBolusAgentPhase             = "00189344"
	CTDIvol                             = "00189345"
	CTDIPhantomTypeCodeSequence         = "00189346"
	CalciumScoringMassFactorPatient     = "00189351"
	CalciumScoringMassFactorDevice      = "00189352"
	EnergyWeightingFactor               = "00189353"
	CTAdditionalXRaySourceSequence      = "00189360"
	ProjectionPixelCalibrationSequence  = "00189401"
	DistanceSourceToIsocenter           = "00189402"
	DistanceObjectToTableTop            = "00189403"
	ObjectPixelSpacingInCenterOfBeam    = "00189404"
	PositionerPositionSequence          = "00189405"
	TablePositionSequence               = "00189406"
	CollimatorShapeSequence             = "00189407"
	XAXRFFrameCharacteristicsSequence   = "00189412"
	FrameAcquisitionSequence            = "00189417"
	XRayReceptorType                    = "00189420"
	AcquisitionProtocolName             = "00189423"
	AcquisitionProtocolDescription      = "00189424"
	ContrastBolusIngredientOpaque       = "00189425"
	DistanceReceptorPlaneToDetHousing   = "00189426"
	IntensifierActiveShape              = "00189427"
	IntensifierActiveDimensions         = "00189428"
	PhysicalDetectorSize                = "00189429"
	PositionOfIsocenterProjection       = "00189430"
	FieldOfViewSequence                 = "00189432"
	FieldOfViewDescription              = "00189433"
	ExposureControlSensingRegionsSeq    = "00189434"
	ExposureControlSensingRegionShape   = "00189435"
	ExposureControlSensRegionLeftEdge   = "00189436"
	ExposureControlSensRegionRightEdge  = "00189437"
	CenterOfCircExposControlSensRegion  = "00189440"
	RadiusOfCircExposControlSensRegion  = "00189441"
	ColumnAngulationPatient             = "00189447"
	BeamAngle                           = "00189449"
	FrameDetectorParametersSequence     = "00189451"
	CalculatedAnatomyThickness          = "00189452"
	CalibrationSequence                 = "00189455"
	ObjectThicknessSequence             = "00189456"
	PlaneIdentification                 = "00189457"
	FieldOfViewDimensionsInFloat        = "00189461"
	IsocenterReferenceSystemSequence    = "00189462"
	PositionerIsocenterPrimaryAngle     = "00189463"
	PositionerIsocenterSecondaryAngle   = "00189464"
	PositionerIsocenterDetRotAngle      = "00189465"
	TableXPositionToIsocenter           = "00189466"
	TableYPositionToIsocenter           = "00189467"
	TableZPositionToIsocenter           = "00189468"
	TableHorizontalRotationAngle        = "00189469"
	TableHeadTiltAngle                  = "00189470"
	TableCradleTiltAngle                = "00189471"
	FrameDisplayShutterSequence         = "00189472"
	AcquiredImageAreaDoseProduct        = "00189473"
	CArmPositionerTabletopRelationship  = "00189474"
	XRayGeometrySequence                = "00189476"
	IrradiationEventIDSequence          = "00189477"
	XRay3DFrameTypeSequence             = "00189504"
	ContributingSourcesSequence         = "00189506"
	XRay3DAcquisitionSequence           = "00189507"
	PrimaryPositionerScanArc            = "00189508"
	SecondaryPositionerScanArc          = "00189509"
	PrimaryPositionerScanStartAngle     = "00189510"
	SecondaryPositionerScanStartAngle   = "00189511"
	PrimaryPositionerIncrement          = "00189514"
	SecondaryPositionerIncrement        = "00189515"
	StartAcquisitionDateTime            = "00189516"
	EndAcquisitionDateTime              = "00189517"
	ApplicationName                     = "00189524"
	ApplicationVersion                  = "00189525"
	ApplicationManufacturer             = "00189526"
	AlgorithmType                       = "00189527"
	AlgorithmDescription                = "00189528"
	XRay3DReconstructionSequence        = "00189530"
	ReconstructionDescription           = "00189531"
	PerProjectionAcquisitionSequence    = "00189538"
	DiffusionBMatrixSequence            = "00189601"
	DiffusionBValueXX                   = "00189602"
	DiffusionBValueXY                   = "00189603"
	DiffusionBValueXZ                   = "00189604"
	DiffusionBValueYY                   = "00189605"
	DiffusionBValueYZ                   = "00189606"
	DiffusionBValueZZ                   = "00189607"
	DecayCorrectionDateTime             = "00189701"
	StartDensityThreshold               = "00189715"
	TerminationTimeThreshold            = "00189722"
	DetectorGeometry                    = "00189725"
	AxialDetectorDimension              = "00189727"
	PETPositionSequence                 = "00189735"
	NumberOfIterations                  = "00189739"
	NumberOfSubsets                     = "00189740"
	PETFrameTypeSequence                = "00189751"
	ReconstructionType                  = "00189756"
	DecayCorrected                      = "00189758"
	AttenuationCorrected                = "00189759"
	ScatterCorrected                    = "00189760"
	DeadTimeCorrected                   = "00189761"
	GantryMotionCorrected               = "00189762"
	PatientMotionCorrected              = "00189763"
	RandomsCorrected                    = "00189765"
	SensitivityCalibrated               = "00189767"
	DepthsOfFocus                       = "00189801"
	ExclusionStartDatetime              = "00189804"
	ExclusionDuration                   = "00189805"
	ImageDataTypeSequence               = "00189807"
	DataType                            = "00189808"
	AliasedDataType                     = "0018980B"
	ContributingEquipmentSequence       = "0018A001"
	ContributionDateTime                = "0018A002"
	ContributionDescription             = "0018A003"
	NumberOfCellsIInDetector            = "00191002"
	CellNumberAtTheta                   = "00191003"
	CellSpacing                         = "00191004"
	HorizFrameOfRef                     = "0019100F"
	SeriesContrast                      = "00191011"
	LastPseq                            = "00191012"
	StartNumberForBaseline              = "00191013"
	EndNumberForBaseline                = "00191014"
	StartNumberForEnhancedScans         = "00191015"
	EndNumberForEnhancedScans           = "00191016"
	SeriesPlane                         = "00191017"
	FirstScanRas                        = "00191018"
	FirstScanLocation                   = "00191019"
	LastScanRas                         = "0019101A"
	LastScanLoc                         = "0019101B"
	DisplayFieldOfView                  = "0019101E"
	MidScanTime                         = "00191024"
	MidScanFlag                         = "00191025"
	DegreesOfAzimuth                    = "00191026"
	GantryPeriod                        = "00191027"
	XRayOnPosition                      = "0019102A"
	XRayOffPosition                     = "0019102B"
	NumberOfTriggers                    = "0019102C"
	AngleOfFirstView                    = "0019102E"
	TriggerFrequency                    = "0019102F"
	ScanFOVType                         = "00191039"
	StatReconFlag                       = "00191040"
	ComputeType                         = "00191041"
	SegmentNumber                       = "00191042"
	TotalSegmentsRequested              = "00191043"
	InterscanDelay                      = "00191044"
	ViewCompressionFactor               = "00191047"
	TotalNoOfRefChannels                = "0019104A"
	DataSizeForScanData                 = "0019104B"
	ReconPostProcflag                   = "00191052"
	CTWaterNumber                       = "00191057"
	CTBoneNumber                        = "00191058"
	NumberOfChannels                    = "0019105E"
	IncrementBetweenChannels            = "0019105F"
	StartingView                        = "00191060"
	NumberOfViews                       = "00191061"
	IncrementBetweenViews               = "00191062"
	DependantOnNoViewsProcessed         = "0019106A"
	FieldOfViewInDetectorCells          = "0019106B"
	ValueOfBackProjectionButton         = "00191070"
	SetIfFatqEstimatesWereUsed          = "00191071"
	ZChanAvgOverViews                   = "00191072"
	AvgOfLeftRefChansOverViews          = "00191073"
	MaxLeftChanOverViews                = "00191074"
	AvgOfRightRefChansOverViews         = "00191075"
	MaxRightChanOverViews               = "00191076"
	SecondEcho                          = "0019107D"
	NumberOfEchoes                      = "0019107E"
	TableDelta                          = "0019107F"
	Contiguous                          = "00191081"
	PeakSAR                             = "00191084"
	MonitorSAR                          = "00191085"
	CardiacRepetitionTime               = "00191087"
	ImagesPerCardiacCycle               = "00191088"
	ActualReceiveGainAnalog             = "0019108A"
	ActualReceiveGainDigital            = "0019108B"
	DelayAfterTrigger                   = "0019108D"
	Swappf                              = "0019108F"
	PauseInterval                       = "00191090"
	PulseTime                           = "00191091"
	SliceOffsetOnFreqAxis               = "00191092"
	CenterFrequency                     = "00191093"
	TransmitGain                        = "00191094"
	AnalogReceiverGain                  = "00191095"
	DigitalReceiverGain                 = "00191096"
	BitmapDefiningCVs                   = "00191097"
	CenterFreqMethod                    = "00191098"
	PulseSeqMode                        = "0019109B"
	PulseSeqName                        = "0019109C"
	PulseSeqDate                        = "0019109D"
	InternalPulseSeqName                = "0019109E"
	TransmittingCoil                    = "0019109F"
	SurfaceCoilType                     = "001910A0"
	ExtremityCoilFlag                   = "001910A1"
	RawDataRunNumber                    = "001910A2"
	CalibratedFieldStrength             = "001910A3"
	SATFatWaterBone                     = "001910A4"
	ReceiveBandwidth                    = "001910A5"
	UserData01                          = "001910A7"
	UserData02                          = "001910A8"
	UserData03                          = "001910A9"
	UserData04                          = "001910AA"
	UserData05                          = "001910AB"
	UserData06                          = "001910AC"
	UserData07                          = "001910AD"
	UserData08                          = "001910AE"
	UserData09                          = "001910AF"
	UserData10                          = "001910B0"
	UserData11                          = "001910B1"
	UserData12                          = "001910B2"
	UserData13                          = "001910B3"
	UserData14                          = "001910B4"
	UserData15                          = "001910B5"
	UserData16                          = "001910B6"
	UserData17                          = "001910B7"
	UserData18                          = "001910B8"
	UserData19                          = "001910B9"
	UserData20                          = "001910BA"
	UserData21                          = "001910BB"
	UserData22                          = "001910BC"
	UserData23                          = "001910BD"
	ProjectionAngle                     = "001910BE"
	SaturationPlanes                    = "001910C0"
	SurfaceCoilIntensity                = "001910C1"
	SATLocationR                        = "001910C2"
	SATLocationL                        = "001910C3"
	SATLocationA                        = "001910C4"
	SATLocationP                        = "001910C5"
	SATLocationH                        = "001910C6"
	SATLocationF                        = "001910C7"
	SATThicknessRL                      = "001910C8"
	SATThicknessAP                      = "001910C9"
	SATThicknessHF                      = "001910CA"
	PrescribedFlowAxis                  = "001910CB"
	VelocityEncoding                    = "001910CC"
	ThicknessDisclaimer                 = "001910CD"
	PrescanType                         = "001910CE"
	PrescanStatus                       = "001910CF"
	RawDataType                         = "001910D0"
	ProjectionAlgorithm                 = "001910D2"
	FractionalEcho                      = "001910D5"
	PrepPulse                           = "001910D6"
	CardiacPhases                       = "001910D7"
	VariableEchoflag                    = "001910D8"
	ConcatenatedSAT                     = "001910D9"
	ReferenceChannelUsed                = "001910DA"
	BackProjectorCoefficient            = "001910DB"
	PrimarySpeedCorrectionUsed          = "001910DC"
	OverrangeCorrectionUsed             = "001910DD"
	DynamicZAlphaValue                  = "001910DE"
	UserData                            = "001910DF"
	VelocityEncodeScale                 = "001910E2"
	FastPhases                          = "001910F2"
	TransmissionGain                    = "001910F9"
	RelationshipGroupLength             = "00200000"
	StudyInstanceUID                    = "0020000D"
	SeriesInstanceUID                   = "0020000E"
	StudyID                             = "00200010"
	SeriesNumber                        = "00200011"
	AcquisitionNumber                   = "00200012"
	InstanceNumber                      = "00200013"
	IsotopeNumber                       = "00200014"
	PhaseNumber                         = "00200015"
	IntervalNumber                      = "00200016"
	TimeSlotNumber                      = "00200017"
	AngleNumber                         = "00200018"
	ItemNumber                          = "00200019"
	PatientOrientation                  = "00200020"
	OverlayNumber                       = "00200022"
	CurveNumber                         = "00200024"
	LookupTableNumber                   = "00200026"
	ImagePosition                       = "00200030"
	ImagePositionPatient                = "00200032"
	ImageOrientation                    = "00200035"
	ImageOrientationPatient             = "00200037"
	Location                            = "00200050"
	FrameOfReferenceUID                 = "00200052"
	Laterality                          = "00200060"
	ImageLaterality                     = "00200062"
	ImageGeometryType                   = "00200070"
	MaskingImage                        = "00200080"
	TemporalPositionIdentifier          = "00200100"
	NumberOfTemporalPositions           = "00200105"
	TemporalResolution                  = "00200110"
	SynchronizationFrameOfReferenceUID  = "00200200"
	SeriesInStudy                       = "00201000"
	AcquisitionsInSeries                = "00201001"
	ImagesInAcquisition                 = "00201002"
	ImagesInSeries                      = "00201003"
	AcquisitionsInStudy                 = "00201004"
	ImagesInStudy                       = "00201005"
	Reference                           = "00201020"
	PositionReferenceIndicator          = "00201040"
	SliceLocation                       = "00201041"
	OtherStudyNumbers                   = "00201070"
	NumberOfPatientRelatedStudies       = "00201200"
	NumberOfPatientRelatedSeries        = "00201202"
	NumberOfPatientRelatedInstances     = "00201204"
	NumberOfStudyRelatedSeries          = "00201206"
	NumberOfStudyRelatedInstances       = "00201208"
	NumberOfSeriesRelatedInstances      = "00201209"
	SourceImageIDs                      = "002031xx"
	ModifyingDeviceID                   = "00203401"
	ModifiedImageID                     = "00203402"
	ModifiedImageDate                   = "00203403"
	ModifyingDeviceManufacturer         = "00203404"
	ModifiedImageTime                   = "00203405"
	ModifiedImageDescription            = "00203406"
	ImageComments                       = "00204000"
	OriginalImageIdentification         = "00205000"
	OriginalImageIdentNomenclature      = "00205002"
	StackID                             = "00209056"
	InStackPositionNumber               = "00209057"
	FrameAnatomySequence                = "00209071"
	FrameLaterality                     = "00209072"
	FrameContentSequence                = "00209111"
	PlanePositionSequence               = "00209113"
	PlaneOrientationSequence            = "00209116"
	TemporalPositionIndex               = "00209128"
	TriggerDelayTime                    = "00209153"
	FrameAcquisitionNumber              = "00209156"
	DimensionIndexValues                = "00209157"
	FrameComments                       = "00209158"
	ConcatenationUID                    = "00209161"
	InConcatenationNumber               = "00209162"
	InConcatenationTotalNumber          = "00209163"
	DimensionOrganizationUID            = "00209164"
	DimensionIndexPointer               = "00209165"
	FunctionalGroupPointer              = "00209167"
	DimensionIndexPrivateCreator        = "00209213"
	DimensionOrganizationSequence       = "00209221"
	DimensionIndexSequence              = "00209222"
	ConcatenationFrameOffsetNumber      = "00209228"
	FunctionalGroupPrivateCreator       = "00209238"
	NominalPercentageOfCardiacPhase     = "00209241"
	NominalPercentOfRespiratoryPhase    = "00209245"
	StartingRespiratoryAmplitude        = "00209246"
	StartingRespiratoryPhase            = "00209247"
	EndingRespiratoryAmplitude          = "00209248"
	EndingRespiratoryPhase              = "00209249"
	RespiratoryTriggerType              = "00209250"
	RRIntervalTimeNominal               = "00209251"
	ActualCardiacTriggerDelayTime       = "00209252"
	RespiratorySynchronizationSequence  = "00209253"
	RespiratoryIntervalTime             = "00209254"
	NominalRespiratoryTriggerDelayTime  = "00209255"
	RespiratoryTriggerDelayThreshold    = "00209256"
	ActualRespiratoryTriggerDelayTime   = "00209257"
	ImagePositionVolume                 = "00209301"
	ImageOrientationVolume              = "00209302"
	ApexPosition                        = "00209308"
	DimensionDescriptionLabel           = "00209421"
	PatientOrientationInFrameSequence   = "00209450"
	FrameLabel                          = "00209453"
	AcquisitionIndex                    = "00209518"
	ContributingSOPInstancesRefSeq      = "00209529"
	ReconstructionIndex                 = "00209536"
	SeriesFromWhichPrescribed           = "00211003"
	SeriesRecordChecksum                = "00211007"
	AcqreconRecordChecksum              = "00211019"
	TableStartLocation                  = "00211020"
	ImageFromWhichPrescribed            = "00211036"
	ScreenFormat                        = "00211037"
	AnatomicalReferenceForScout         = "0021104A"
	LocationsInAcquisition              = "0021104F"
	GraphicallyPrescribed               = "00211050"
	RotationFromSourceXRot              = "00211051"
	RotationFromSourceYRot              = "00211052"
	RotationFromSourceZRot              = "00211053"
	IntegerSlop                         = "00211056"
	FloatSlop                           = "0021105F"
	AutoWindowLevelAlpha                = "00211081"
	AutoWindowLevelBeta                 = "00211082"
	AutoWindowLevelWindow               = "00211083"
	ToWindowLevelLevel                  = "00211084"
	TubeFocalSpotPosition               = "00211090"
	BiopsyPosition                      = "00211091"
	BiopsyTLocation                     = "00211092"
	BiopsyRefLocation                   = "00211093"
	LightPathFilterPassThroughWavelen   = "00220001"
	LightPathFilterPassBand             = "00220002"
	ImagePathFilterPassThroughWavelen   = "00220003"
	ImagePathFilterPassBand             = "00220004"
	PatientEyeMovementCommanded         = "00220005"
	PatientEyeMovementCommandCodeSeq    = "00220006"
	SphericalLensPower                  = "00220007"
	CylinderLensPower                   = "00220008"
	CylinderAxis                        = "00220009"
	EmmetropicMagnification             = "0022000A"
	IntraOcularPressure                 = "0022000B"
	HorizontalFieldOfView               = "0022000C"
	PupilDilated                        = "0022000D"
	DegreeOfDilation                    = "0022000E"
	StereoBaselineAngle                 = "00220010"
	StereoBaselineDisplacement          = "00220011"
	StereoHorizontalPixelOffset         = "00220012"
	StereoVerticalPixelOffset           = "00220013"
	StereoRotation                      = "00220014"
	AcquisitionDeviceTypeCodeSequence   = "00220015"
	IlluminationTypeCodeSequence        = "00220016"
	LightPathFilterTypeStackCodeSeq     = "00220017"
	ImagePathFilterTypeStackCodeSeq     = "00220018"
	LensesCodeSequence                  = "00220019"
	ChannelDescriptionCodeSequence      = "0022001A"
	RefractiveStateSequence             = "0022001B"
	MydriaticAgentCodeSequence          = "0022001C"
	RelativeImagePositionCodeSequence   = "0022001D"
	StereoPairsSequence                 = "00220020"
	LeftImageSequence                   = "00220021"
	RightImageSequence                  = "00220022"
	AxialLengthOfTheEye                 = "00220030"
	OphthalmicFrameLocationSequence     = "00220031"
	ReferenceCoordinates                = "00220032"
	DepthSpatialResolution              = "00220035"
	MaximumDepthDistortion              = "00220036"
	AlongScanSpatialResolution          = "00220037"
	MaximumAlongScanDistortion          = "00220038"
	OphthalmicImageOrientation          = "00220039"
	DepthOfTransverseImage              = "00220041"
	MydriaticAgentConcUnitsSeq          = "00220042"
	AcrossScanSpatialResolution         = "00220048"
	MaximumAcrossScanDistortion         = "00220049"
	MydriaticAgentConcentration         = "0022004E"
	IlluminationWaveLength              = "00220055"
	IlluminationPower                   = "00220056"
	IlluminationBandwidth               = "00220057"
	MydriaticAgentSequence              = "00220058"
	NumberOfSeriesInStudy               = "00231001"
	NumberOfUnarchivedSeries            = "00231002"
	ReferenceImageField                 = "00231010"
	SummaryImage                        = "00231050"
	StartTimeSecsInFirstAxial           = "00231070"
	NoofUpdatesToHeader                 = "00231074"
	IndicatesIfTheStudyHasCompleteInfo  = "0023107D"
	LastPulseSequenceUsed               = "00251006"
	LandmarkCounter                     = "00251010"
	NumberOfAcquisitions                = "00251011"
	IndicatesNoofUpdatesToHeader        = "00251014"
	SeriesCompleteFlag                  = "00251017"
	NumberOfImagesArchived              = "00251018"
	LastImageNumberUsed                 = "00251019"
	PrimaryReceiverSuiteAndHost         = "0025101A"
	ImageArchiveFlag                    = "00271006"
	ScoutType                           = "00271010"
	VmaMamp                             = "0027101C"
	VmaPhase                            = "0027101D"
	VmaMod                              = "0027101E"
	VmaClip                             = "0027101F"
	SmartScanOnOffFlag                  = "00271020"
	ForeignImageRevision                = "00271030"
	ImagingMode                         = "00271031"
	PulseSequence                       = "00271032"
	ImagingOptions                      = "00271033"
	PlaneType                           = "00271035"
	ObliquePlane                        = "00271036"
	RASLetterOfImageLocation            = "00271040"
	ImageLocation                       = "00271041"
	CenterRCoordOfPlaneImage            = "00271042"
	CenterACoordOfPlaneImage            = "00271043"
	CenterSCoordOfPlaneImage            = "00271044"
	NormalRCoord                        = "00271045"
	NormalACoord                        = "00271046"
	NormalSCoord                        = "00271047"
	RCoordOfTopRightCorner              = "00271048"
	ACoordOfTopRightCorner              = "00271049"
	SCoordOfTopRightCorner              = "0027104A"
	RCoordOfBottomRightCorner           = "0027104B"
	ACoordOfBottomRightCorner           = "0027104C"
	SCoordOfBottomRightCorner           = "0027104D"
	TableEndLocation                    = "00271051"
	RASLetterForSideOfImage             = "00271052"
	RASLetterForAnteriorPosterior       = "00271053"
	RASLetterForScoutStartLoc           = "00271054"
	RASLetterForScoutEndLoc             = "00271055"
	ImageDimensionX                     = "00271060"
	ImageDimensionY                     = "00271061"
	NumberOfExcitations                 = "00271062"
	ImagePresentationGroupLength        = "00280000"
	SamplesPerPixel                     = "00280002"
	SamplesPerPixelUsed                 = "00280003"
	PhotometricInterpretation           = "00280004"
	ImageDimensions                     = "00280005"
	PlanarConfiguration                 = "00280006"
	NumberOfFrames                      = "00280008"
	FrameIncrementPointer               = "00280009"
	FrameDimensionPointer               = "0028000A"
	Rows                                = "00280010"
	Columns                             = "00280011"
	Planes                              = "00280012"
	UltrasoundColorDataPresent          = "00280014"
	PixelSpacing                        = "00280030"
	ZoomFactor                          = "00280031"
	ZoomCenter                          = "00280032"
	PixelAspectRatio                    = "00280034"
	ImageFormat                         = "00280040"
	ManipulatedImage                    = "00280050"
	CorrectedImage                      = "00280051"
	CompressionRecognitionCode          = "0028005F"
	CompressionCode                     = "00280060"
	CompressionOriginator               = "00280061"
	CompressionLabel                    = "00280062"
	CompressionDescription              = "00280063"
	CompressionSequence                 = "00280065"
	CompressionStepPointers             = "00280066"
	RepeatInterval                      = "00280068"
	BitsGrouped                         = "00280069"
	PerimeterTable                      = "00280070"
	PerimeterValue                      = "00280071"
	PredictorRows                       = "00280080"
	PredictorColumns                    = "00280081"
	PredictorConstants                  = "00280082"
	BlockedPixels                       = "00280090"
	BlockRows                           = "00280091"
	BlockColumns                        = "00280092"
	RowOverlap                          = "00280093"
	ColumnOverlap                       = "00280094"
	BitsAllocated                       = "00280100"
	BitsStored                          = "00280101"
	HighBit                             = "00280102"
	PixelRepresentation                 = "00280103"
	SmallestValidPixelValue             = "00280104"
	LargestValidPixelValue              = "00280105"
	SmallestImagePixelValue             = "00280106"
	LargestImagePixelValue              = "00280107"
	SmallestPixelValueInSeries          = "00280108"
	LargestPixelValueInSeries           = "00280109"
	SmallestImagePixelValueInPlane      = "00280110"
	LargestImagePixelValueInPlane       = "00280111"
	PixelPaddingValue                   = "00280120"
	PixelPaddingRangeLimit              = "00280121"
	QualityControlImage                 = "00280300"
	BurnedInAnnotation                  = "00280301"
	TransformLabel                      = "00280400"
	TransformVersionNumber              = "00280401"
	NumberOfTransformSteps              = "00280402"
	SequenceOfCompressedData            = "00280403"
	DetailsOfCoefficients               = "00280404"
	CoefficientCoding                   = "002804x2"
	CoefficientCodingPointers           = "002804x3"
	DCTLabel                            = "00280700"
	DataBlockDescription                = "00280701"
	DataBlock                           = "00280702"
	NormalizationFactorFormat           = "00280710"
	ZonalMapNumberFormat                = "00280720"
	ZonalMapLocation                    = "00280721"
	ZonalMapFormat                      = "00280722"
	AdaptiveMapFormat                   = "00280730"
	CodeNumberFormat                    = "00280740"
	CodeLabel                           = "002808x0"
	NumberOfTables                      = "002808x2"
	CodeTableLocation                   = "002808x3"
	BitsForCodeWord                     = "002808x4"
	ImageDataLocation                   = "002808x8"
	PixelSpacingCalibrationType         = "00280A02"
	PixelSpacingCalibrationDescription  = "00280A04"
	PixelIntensityRelationship          = "00281040"
	PixelIntensityRelationshipSign      = "00281041"
	WindowCenter                        = "00281050"
	WindowWidth                         = "00281051"
	RescaleIntercept                    = "00281052"
	RescaleSlope                        = "00281053"
	RescaleType                         = "00281054"
	WindowCenterAndWidthExplanation     = "00281055"
	VOILUTFunction                      = "00281056"
	GrayScale                           = "00281080"
	RecommendedViewingMode              = "00281090"
	GrayLookupTableDescriptor           = "00281100"
	RedPaletteColorTableDescriptor      = "00281101"
	GreenPaletteColorTableDescriptor    = "00281102"
	BluePaletteColorTableDescriptor     = "00281103"
	LargeRedPaletteColorTableDescr      = "00281111"
	LargeGreenPaletteColorTableDescr    = "00281112"
	LargeBluePaletteColorTableDescr     = "00281113"
	PaletteColorTableUID                = "00281199"
	GrayLookupTableData                 = "00281200"
	RedPaletteColorTableData            = "00281201"
	GreenPaletteColorTableData          = "00281202"
	BluePaletteColorTableData           = "00281203"
	LargeRedPaletteColorTableData       = "00281211"
	LargeGreenPaletteColorTableData     = "00281212"
	LargeBluePaletteColorTableData      = "00281213"
	LargePaletteColorLookupTableUID     = "00281214"
	SegmentedRedColorTableData          = "00281221"
	SegmentedGreenColorTableData        = "00281222"
	SegmentedBlueColorTableData         = "00281223"
	BreastImplantPresent                = "00281300"
	PartialView                         = "00281350"
	PartialViewDescription              = "00281351"
	PartialViewCodeSequence             = "00281352"
	SpatialLocationsPreserved           = "0028135A"
	DataPathAssignment                  = "00281402"
	BlendingLUT1Sequence                = "00281404"
	BlendingWeightConstant              = "00281406"
	BlendingLookupTableData             = "00281408"
	BlendingLUT2Sequence                = "0028140C"
	DataPathID                          = "0028140E"
	RGBLUTTransferFunction              = "0028140F"
	AlphaLUTTransferFunction            = "00281410"
	ICCProfile                          = "00282000"
	LossyImageCompressionRatio          = "00282112"
	LossyImageCompressionMethod         = "00282114"
	ModalityLUTSequence                 = "00283000"
	LUTDescriptor                       = "00283002"
	LUTExplanation                      = "00283003"
	ModalityLUTType                     = "00283004"
	LUTData                             = "00283006"
	VOILUTSequence                      = "00283010"
	SoftcopyVOILUTSequence              = "00283110"
	ImagePresentationComments           = "00284000"
	BiPlaneAcquisitionSequence          = "00285000"
	RepresentativeFrameNumber           = "00286010"
	FrameNumbersOfInterest              = "00286020"
	FrameOfInterestDescription          = "00286022"
	FrameOfInterestType                 = "00286023"
	MaskPointers                        = "00286030"
	RWavePointer                        = "00286040"
	MaskSubtractionSequence             = "00286100"
	MaskOperation                       = "00286101"
	ApplicableFrameRange                = "00286102"
	MaskFrameNumbers                    = "00286110"
	ContrastFrameAveraging              = "00286112"
	MaskSubPixelShift                   = "00286114"
	TIDOffset                           = "00286120"
	MaskOperationExplanation            = "00286190"
	PixelDataProviderURL                = "00287FE0"
	DataPointRows                       = "00289001"
	DataPointColumns                    = "00289002"
	SignalDomainColumns                 = "00289003"
	LargestMonochromePixelValue         = "00289099"
	DataRepresentation                  = "00289108"
	PixelMeasuresSequence               = "00289110"
	FrameVOILUTSequence                 = "00289132"
	PixelValueTransformationSequence    = "00289145"
	SignalDomainRows                    = "00289235"
	DisplayFilterPercentage             = "00289411"
	FramePixelShiftSequence             = "00289415"
	SubtractionItemID                   = "00289416"
	PixelIntensityRelationshipLUTSeq    = "00289422"
	FramePixelDataPropertiesSequence    = "00289443"
	GeometricalProperties               = "00289444"
	GeometricMaximumDistortion          = "00289445"
	ImageProcessingApplied              = "00289446"
	MaskSelectionMode                   = "00289454"
	LUTFunction                         = "00289474"
	MaskVisibilityPercentage            = "00289478"
	PixelShiftSequence                  = "00289501"
	RegionPixelShiftSequence            = "00289502"
	VerticesOfTheRegion                 = "00289503"
	PixelShiftFrameRange                = "00289506"
	LUTFrameRange                       = "00289507"
	ImageToEquipmentMappingMatrix       = "00289520"
	EquipmentCoordinateSystemID         = "00289537"
	LowerRangeOfPixels1a                = "00291004"
	LowerRangeOfPixels1b                = "00291005"
	LowerRangeOfPixels1c                = "00291006"
	LowerRangeOfPixels1d                = "00291007"
	LowerRangeOfPixels1e                = "00291008"
	LowerRangeOfPixels1f                = "00291009"
	LowerRangeOfPixels1g                = "0029100A"
	LowerRangeOfPixels1h                = "00291015"
	LowerRangeOfPixels1i                = "00291016"
	LowerRangeOfPixels2                 = "00291017"
	UpperRangeOfPixels2                 = "00291018"
	LenOfTotHdrInBytes                  = "0029101A"
	VersionOfTheHdrStruct               = "00291026"
	AdvantageCompOverflow               = "00291034"
	AdvantageCompUnderflow              = "00291035"
	StudyGroupLength                    = "00320000"
	StudyStatusID                       = "0032000A"
	StudyPriorityID                     = "0032000C"
	StudyIDIssuer                       = "00320012"
	StudyVerifiedDate                   = "00320032"
	StudyVerifiedTime                   = "00320033"
	StudyReadDate                       = "00320034"
	StudyReadTime                       = "00320035"
	ScheduledStudyStartDate             = "00321000"
	ScheduledStudyStartTime             = "00321001"
	ScheduledStudyStopDate              = "00321010"
	ScheduledStudyStopTime              = "00321011"
	ScheduledStudyLocation              = "00321020"
	ScheduledStudyLocationAETitle       = "00321021"
	ReasonForStudy                      = "00321030"
	RequestingPhysicianIDSequence       = "00321031"
	RequestingPhysician                 = "00321032"
	RequestingService                   = "00321033"
	StudyArrivalDate                    = "00321040"
	StudyArrivalTime                    = "00321041"
	StudyCompletionDate                 = "00321050"
	StudyCompletionTime                 = "00321051"
	StudyComponentStatusID              = "00321055"
	RequestedProcedureDescription       = "00321060"
	RequestedProcedureCodeSequence      = "00321064"
	RequestedContrastAgent              = "00321070"
	StudyComments                       = "00324000"
	ReferencedPatientAliasSequence      = "00380004"
	VisitStatusID                       = "00380008"
	AdmissionID                         = "00380010"
	IssuerOfAdmissionID                 = "00380011"
	RouteOfAdmissions                   = "00380016"
	ScheduledAdmissionDate              = "0038001A"
	ScheduledAdmissionTime              = "0038001B"
	ScheduledDischargeDate              = "0038001C"
	ScheduledDischargeTime              = "0038001D"
	ScheduledPatientInstitResidence     = "0038001E"
	AdmittingDate                       = "00380020"
	AdmittingTime                       = "00380021"
	DischargeDate                       = "00380030"
	DischargeTime                       = "00380032"
	DischargeDiagnosisDescription       = "00380040"
	DischargeDiagnosisCodeSequence      = "00380044"
	SpecialNeeds                        = "00380050"
	ServiceEpisodeID                    = "00380060"
	IssuerOfServiceEpisodeID            = "00380061"
	ServiceEpisodeDescription           = "00380062"
	PertinentDocumentsSequence          = "00380100"
	CurrentPatientLocation              = "00380300"
	PatientInstitutionResidence         = "00380400"
	PatientState                        = "00380500"
	PatientClinicalTrialParticipSeq     = "00380502"
	VisitComments                       = "00384000"
	WaveformOriginality                 = "003A0004"
	NumberOfWaveformChannels            = "003A0005"
	NumberOfWaveformSamples             = "003A0010"
	SamplingFrequency                   = "003A001A"
	MultiplexGroupLabel                 = "003A0020"
	ChannelDefinitionSequence           = "003A0200"
	WaveformChannelNumber               = "003A0202"
	ChannelLabel                        = "003A0203"
	ChannelStatus                       = "003A0205"
	ChannelSourceSequence               = "003A0208"
	ChannelSourceModifiersSequence      = "003A0209"
	SourceWaveformSequence              = "003A020A"
	ChannelDerivationDescription        = "003A020C"
	ChannelSensitivity                  = "003A0210"
	ChannelSensitivityUnitsSequence     = "003A0211"
	ChannelSensitivityCorrectionFactor  = "003A0212"
	ChannelBaseline                     = "003A0213"
	ChannelTimeSkew                     = "003A0214"
	ChannelSampleSkew                   = "003A0215"
	ChannelOffset                       = "003A0218"
	WaveformBitsStored                  = "003A021A"
	FilterLowFrequency                  = "003A0220"
	FilterHighFrequency                 = "003A0221"
	NotchFilterFrequency                = "003A0222"
	NotchFilterBandwidth                = "003A0223"
	WaveformDataDisplayScale            = "003A0230"
	WaveformDisplayBkgCIELabValue       = "003A0231"
	WaveformPresentationGroupSequence   = "003A0240"
	PresentationGroupNumber             = "003A0241"
	ChannelDisplaySequence              = "003A0242"
	ChannelRecommendDisplayCIELabValue  = "003A0244"
	ChannelPosition                     = "003A0245"
	DisplayShadingFlag                  = "003A0246"
	FractionalChannelDisplayScale       = "003A0247"
	AbsoluteChannelDisplayScale         = "003A0248"
	MultiplexAudioChannelsDescrCodeSeq  = "003A0300"
	ChannelIdentificationCode           = "003A0301"
	ChannelMode                         = "003A0302"
	ScheduledStationAETitle             = "00400001"
	ScheduledProcedureStepStartDate     = "00400002"
	ScheduledProcedureStepStartTime     = "00400003"
	ScheduledProcedureStepEndDate       = "00400004"
	ScheduledProcedureStepEndTime       = "00400005"
	ScheduledPerformingPhysiciansName   = "00400006"
	ScheduledProcedureStepDescription   = "00400007"
	ScheduledProtocolCodeSequence       = "00400008"
	ScheduledProcedureStepID            = "00400009"
	StageCodeSequence                   = "0040000A"
	ScheduledPerformingPhysicianIDSeq   = "0040000B"
	ScheduledStationName                = "00400010"
	ScheduledProcedureStepLocation      = "00400011"
	PreMedication                       = "00400012"
	ScheduledProcedureStepStatus        = "00400020"
	LocalNamespaceEntityID              = "00400031"
	UniversalEntityID                   = "00400032"
	UniversalEntityIDType               = "00400033"
	IdentifierTypeCode                  = "00400035"
	AssigningFacilitySequence           = "00400036"
	ScheduledProcedureStepSequence      = "00400100"
	ReferencedNonImageCompositeSOPSeq   = "00400220"
	PerformedStationAETitle             = "00400241"
	PerformedStationName                = "00400242"
	PerformedLocation                   = "00400243"
	PerformedProcedureStepStartDate     = "00400244"
	PerformedProcedureStepStartTime     = "00400245"
	PerformedProcedureStepEndDate       = "00400250"
	PerformedProcedureStepEndTime       = "00400251"
	PerformedProcedureStepStatus        = "00400252"
	PerformedProcedureStepID            = "00400253"
	PerformedProcedureStepDescription   = "00400254"
	PerformedProcedureTypeDescription   = "00400255"
	PerformedProtocolCodeSequence       = "00400260"
	PerformedProtocolType               = "00400261"
	ScheduledStepAttributesSequence     = "00400270"
	RequestAttributesSequence           = "00400275"
	CommentsOnPerformedProcedureStep    = "00400280"
	ProcStepDiscontinueReasonCodeSeq    = "00400281"
	QuantitySequence                    = "00400293"
	Quantity                            = "00400294"
	MeasuringUnitsSequence              = "00400295"
	BillingItemSequence                 = "00400296"
	TotalTimeOfFluoroscopy              = "00400300"
	TotalNumberOfExposures              = "00400301"
	EntranceDose                        = "00400302"
	ExposedArea                         = "00400303"
	DistanceSourceToEntrance            = "00400306"
	DistanceSourceToSupport             = "00400307"
	ExposureDoseSequence                = "0040030E"
	CommentsOnRadiationDose             = "00400310"
	XRayOutput                          = "00400312"
	HalfValueLayer                      = "00400314"
	OrganDose                           = "00400316"
	OrganExposed                        = "00400318"
	BillingProcedureStepSequence        = "00400320"
	FilmConsumptionSequence             = "00400321"
	BillingSuppliesAndDevicesSequence   = "00400324"
	PerformedSeriesSequence             = "00400340"
	CommentsOnScheduledProcedureStep    = "00400400"
	ProtocolContextSequence             = "00400440"
	ContentItemModifierSequence         = "00400441"
	SpecimenAccessionNumber             = "0040050A"
	ContainerIdentifier                 = "00400512"
	ContainerDescription                = "0040051A"
	SpecimenSequence                    = "00400550"
	SpecimenIdentifier                  = "00400551"
	SpecimenDescriptionSequenceTrial    = "00400552"
	SpecimenDescriptionTrial            = "00400553"
	SpecimenUID                         = "00400554"
	AcquisitionContextSequence          = "00400555"
	AcquisitionContextDescription       = "00400556"
	SpecimenTypeCodeSequence            = "0040059A"
	SpecimenShortDescription            = "00400600"
	SlideIdentifier                     = "004006FA"
	ImageCenterPointCoordinatesSeq      = "0040071A"
	XOffsetInSlideCoordinateSystem      = "0040072A"
	YOffsetInSlideCoordinateSystem      = "0040073A"
	ZOffsetInSlideCoordinateSystem      = "0040074A"
	PixelSpacingSequence                = "004008D8"
	CoordinateSystemAxisCodeSequence    = "004008DA"
	MeasurementUnitsCodeSequence        = "004008EA"
	VitalStainCodeSequenceTrial         = "004009F8"
	RequestedProcedureID                = "00401001"
	ReasonForRequestedProcedure         = "00401002"
	RequestedProcedurePriority          = "00401003"
	PatientTransportArrangements        = "00401004"
	RequestedProcedureLocation          = "00401005"
	PlacerOrderNumberProcedure          = "00401006"
	FillerOrderNumberProcedure          = "00401007"
	ConfidentialityCode                 = "00401008"
	ReportingPriority                   = "00401009"
	ReasonForRequestedProcedureCodeSeq  = "0040100A"
	NamesOfIntendedRecipientsOfResults  = "00401010"
	IntendedRecipientsOfResultsIDSeq    = "00401011"
	PersonIdentificationCodeSequence    = "00401101"
	PersonAddress                       = "00401102"
	PersonTelephoneNumbers              = "00401103"
	RequestedProcedureComments          = "00401400"
	ReasonForImagingServiceRequest      = "00402001"
	IssueDateOfImagingServiceRequest    = "00402004"
	IssueTimeOfImagingServiceRequest    = "00402005"
	PlacerOrderNumImagingServiceReq     = "00402006"
	FillerOrderNumImagingServiceReq     = "00402007"
	OrderEnteredBy                      = "00402008"
	OrderEntererLocation                = "00402009"
	OrderCallbackPhoneNumber            = "00402010"
	ImagingServiceRequestComments       = "00402400"
	ConfidentialityOnPatientDataDescr   = "00403001"
	GenPurposeScheduledProcStepStatus   = "00404001"
	GenPurposePerformedProcStepStatus   = "00404002"
	GenPurposeSchedProcStepPriority     = "00404003"
	SchedProcessingApplicationsCodeSeq  = "00404004"
	SchedProcedureStepStartDateAndTime  = "00404005"
	MultipleCopiesFlag                  = "00404006"
	PerformedProcessingAppsCodeSeq      = "00404007"
	HumanPerformerCodeSequence          = "00404009"
	SchedProcStepModificationDateTime   = "00404010"
	ExpectedCompletionDateAndTime       = "00404011"
	ResultingGenPurposePerfProcStepSeq  = "00404015"
	RefGenPurposeSchedProcStepSeq       = "00404016"
	ScheduledWorkitemCodeSequence       = "00404018"
	PerformedWorkitemCodeSequence       = "00404019"
	InputAvailabilityFlag               = "00404020"
	InputInformationSequence            = "00404021"
	RelevantInformationSequence         = "00404022"
	RefGenPurSchedProcStepTransUID      = "00404023"
	ScheduledStationNameCodeSequence    = "00404025"
	ScheduledStationClassCodeSequence   = "00404026"
	SchedStationGeographicLocCodeSeq    = "00404027"
	PerformedStationNameCodeSequence    = "00404028"
	PerformedStationClassCodeSequence   = "00404029"
	PerformedStationGeogLocCodeSeq      = "00404030"
	RequestedSubsequentWorkItemCodeSeq  = "00404031"
	NonDICOMOutputCodeSequence          = "00404032"
	OutputInformationSequence           = "00404033"
	ScheduledHumanPerformersSequence    = "00404034"
	ActualHumanPerformersSequence       = "00404035"
	HumanPerformersOrganization         = "00404036"
	HumanPerformerName                  = "00404037"
	RawDataHandling                     = "00404040"
	EntranceDoseInMilliGy               = "00408302"
	RefImageRealWorldValueMappingSeq    = "00409094"
	RealWorldValueMappingSequence       = "00409096"
	PixelValueMappingCodeSequence       = "00409098"
	LUTLabel                            = "00409210"
	RealWorldValueLastValueMapped       = "00409211"
	RealWorldValueLUTData               = "00409212"
	RealWorldValueFirstValueMapped      = "00409216"
	RealWorldValueIntercept             = "00409224"
	RealWorldValueSlope                 = "00409225"
	RelationshipType                    = "0040A010"
	VerifyingOrganization               = "0040A027"
	VerificationDateTime                = "0040A030"
	ObservationDateTime                 = "0040A032"
	ValueType                           = "0040A040"
	ConceptNameCodeSequence             = "0040A043"
	ContinuityOfContent                 = "0040A050"
	VerifyingObserverSequence           = "0040A073"
	VerifyingObserverName               = "0040A075"
	AuthorObserverSequence              = "0040A078"
	ParticipantSequence                 = "0040A07A"
	CustodialOrganizationSequence       = "0040A07C"
	ParticipationType                   = "0040A080"
	ParticipationDateTime               = "0040A082"
	ObserverType                        = "0040A084"
	VerifyingObserverIdentCodeSequence  = "0040A088"
	EquivalentCDADocumentSequence       = "0040A090"
	ReferencedWaveformChannels          = "0040A0B0"
	DateTime                            = "0040A120"
	Date                                = "0040A121"
	Time                                = "0040A122"
	PersonName                          = "0040A123"
	UID                                 = "0040A124"
	TemporalRangeType                   = "0040A130"
	ReferencedSamplePositions           = "0040A132"
	ReferencedFrameNumbers              = "0040A136"
	ReferencedTimeOffsets               = "0040A138"
	ReferencedDateTime                  = "0040A13A"
	TextValue                           = "0040A160"
	ConceptCodeSequence                 = "0040A168"
	PurposeOfReferenceCodeSequence      = "0040A170"
	AnnotationGroupNumber               = "0040A180"
	ModifierCodeSequence                = "0040A195"
	MeasuredValueSequence               = "0040A300"
	NumericValueQualifierCodeSequence   = "0040A301"
	NumericValue                        = "0040A30A"
	AddressTrial                        = "0040A353"
	TelephoneNumberTrial                = "0040A354"
	PredecessorDocumentsSequence        = "0040A360"
	ReferencedRequestSequence           = "0040A370"
	PerformedProcedureCodeSequence      = "0040A372"
	CurrentRequestedProcEvidenceSeq     = "0040A375"
	PertinentOtherEvidenceSequence      = "0040A385"
	HL7StructuredDocumentRefSeq         = "0040A390"
	CompletionFlag                      = "0040A491"
	CompletionFlagDescription           = "0040A492"
	VerificationFlag                    = "0040A493"
	ArchiveRequested                    = "0040A494"
	PreliminaryFlag                     = "0040A496"
	ContentTemplateSequence             = "0040A504"
	IdenticalDocumentsSequence          = "0040A525"
	ContentSequence                     = "0040A730"
	AnnotationSequence                  = "0040B020"
	TemplateIdentifier                  = "0040DB00"
	TemplateVersion                     = "0040DB06"
	TemplateLocalVersion                = "0040DB07"
	TemplateExtensionFlag               = "0040DB0B"
	TemplateExtensionOrganizationUID    = "0040DB0C"
	TemplateExtensionCreatorUID         = "0040DB0D"
	ReferencedContentItemIdentifier     = "0040DB73"
	HL7InstanceIdentifier               = "0040E001"
	HL7DocumentEffectiveTime            = "0040E004"
	HL7DocumentTypeCodeSequence         = "0040E006"
	RetrieveURI                         = "0040E010"
	RetrieveLocationUID                 = "0040E011"
	DocumentTitle                       = "00420010"
	EncapsulatedDocument                = "00420011"
	MIMETypeOfEncapsulatedDocument      = "00420012"
	SourceInstanceSequence              = "00420013"
	ListOfMIMETypes                     = "00420014"
	BitmapOfPrescanOptions              = "00431001"
	GradientOffsetInX                   = "00431002"
	GradientOffsetInY                   = "00431003"
	GradientOffsetInZ                   = "00431004"
	ImgIsOriginalOrUnoriginal           = "00431005"
	NumberOfEPIShots                    = "00431006"
	ViewsPerSegment                     = "00431007"
	RespiratoryRateBpm                  = "00431008"
	RespiratoryTriggerPoint             = "00431009"
	TypeOfReceiverUsed                  = "0043100A"
	PeakRateOfChangeOfGradientField     = "0043100B"
	LimitsInUnitsOfPercent              = "0043100C"
	PSDEstimatedLimit                   = "0043100D"
	PSDEstimatedLimitInTeslaPerSecond   = "0043100E"
	Saravghead                          = "0043100F"
	WindowValue                         = "00431010"
	TotalInputViews                     = "00431011"
	XRayChain                           = "00431012"
	DeconKernelParameters               = "00431013"
	CalibrationParameters               = "00431014"
	TotalOutputViews                    = "00431015"
	NumberOfOverranges                  = "00431016"
	IBHImageScaleFactors                = "00431017"
	BBHCoefficients                     = "00431018"
	NumberOfBBHChainsToBlend            = "00431019"
	StartingChannelNumber               = "0043101A"
	PpscanParameters                    = "0043101B"
	GEImageIntegrity                    = "0043101C"
	LevelValue                          = "0043101D"
	DeltaStartTime                      = "0043101E"
	MaxOverrangesInAView                = "0043101F"
	AvgOverrangesAllViews               = "00431020"
	CorrectedAfterGlowTerms             = "00431021"
	ReferenceChannels                   = "00431025"
	NoViewsRefChansBlocked              = "00431026"
	ScanPitchRatio                      = "00431027"
	UniqueImageIden                     = "00431028"
	HistogramTables                     = "00431029"
	UserDefinedData                     = "0043102A"
	PrivateScanOptions                  = "0043102B"
	EffectiveEchoSpacing                = "0043102C"
	StringSlopField1                    = "0043102D"
	StringSlopField2                    = "0043102E"
	RACordOfTargetReconCenter           = "00431031"
	NegScanspacing                      = "00431033"
	OffsetFrequency                     = "00431034"
	UserUsageTag                        = "00431035"
	UserFillMapMSW                      = "00431036"
	UserFillMapLSW                      = "00431037"
	User2548                            = "00431038"
	SlopInt69                           = "00431039"
	TriggerOnPosition                   = "00431040"
	DegreeOfRotation                    = "00431041"
	DASTriggerSource                    = "00431042"
	DASFpaGain                          = "00431043"
	DASOutputSource                     = "00431044"
	DASAdInput                          = "00431045"
	DASCalMode                          = "00431046"
	DASCalFrequency                     = "00431047"
	DASRegXm                            = "00431048"
	DASAutoZero                         = "00431049"
	StartingChannelOfView               = "0043104A"
	DASXmPattern                        = "0043104B"
	TGGCTriggerMode                     = "0043104C"
	StartScanToXrayOnDelay              = "0043104D"
	DurationOfXrayOn                    = "0043104E"
	SlopInt1017                         = "00431060"
	ScannerStudyEntityUID               = "00431061"
	ScannerStudyID                      = "00431062"
	ScannerTableEntry                   = "0043106f"
	ProductPackageIdentifier            = "00440001"
	SubstanceAdministrationApproval     = "00440002"
	ApprovalStatusFurtherDescription    = "00440003"
	ApprovalStatusDateTime              = "00440004"
	ProductTypeCodeSequence             = "00440007"
	ProductName                         = "00440008"
	ProductDescription                  = "00440009"
	ProductLotIdentifier                = "0044000A"
	ProductExpirationDateTime           = "0044000B"
	SubstanceAdministrationDateTime     = "00440010"
	SubstanceAdministrationNotes        = "00440011"
	SubstanceAdministrationDeviceID     = "00440012"
	ProductParameterSequence            = "00440013"
	SubstanceAdminParameterSeq          = "00440019"
	NumberOfMacroRowsInDetector         = "00451001"
	MacroWidthAtISOCenter               = "00451002"
	DASType                             = "00451003"
	DASGain                             = "00451004"
	DASTemperature                      = "00451005"
	TableDirectionInOrOut               = "00451006"
	ZSmoothingFactor                    = "00451007"
	ViewWeightingMode                   = "00451008"
	SigmaRowNumberWhichRowsWereUsed     = "00451009"
	MinimumDasValueFoundInTheScanData   = "0045100A"
	MaximumOffsetShiftValueUsed         = "0045100B"
	NumberOfViewsShifted                = "0045100C"
	ZTrackingFlag                       = "0045100D"
	MeanZError                          = "0045100E"
	ZTrackingMaximumError               = "0045100F"
	StartingViewForRow2a                = "00451010"
	NumberOfViewsInRow2a                = "00451011"
	StartingViewForRow1a                = "00451012"
	SigmaMode                           = "00451013"
	NumberOfViewsInRow1a                = "00451014"
	StartingViewForRow2b                = "00451015"
	NumberOfViewsInRow2b                = "00451016"
	StartingViewForRow1b                = "00451017"
	NumberOfViewsInRow1b                = "00451018"
	AirFilterCalibrationDate            = "00451019"
	AirFilterCalibrationTime            = "0045101A"
	PhantomCalibrationDate              = "0045101B"
	PhantomCalibrationTime              = "0045101C"
	ZSlopeCalibrationDate               = "0045101D"
	ZSlopeCalibrationTime               = "0045101E"
	CrosstalkCalibrationDate            = "0045101F"
	CrosstalkCalibrationTime            = "00451020"
	IterboneOptionFlag                  = "00451021"
	PeristalticFlagOption               = "00451022"
	LensDescription                     = "00460012"
	RightLensSequence                   = "00460014"
	LeftLensSequence                    = "00460015"
	CylinderSequence                    = "00460018"
	PrismSequence                       = "00460028"
	HorizontalPrismPower                = "00460030"
	HorizontalPrismBase                 = "00460032"
	VerticalPrismPower                  = "00460034"
	VerticalPrismBase                   = "00460036"
	LensSegmentType                     = "00460038"
	OpticalTransmittance                = "00460040"
	ChannelWidth                        = "00460042"
	PupilSize                           = "00460044"
	CornealSize                         = "00460046"
	DistancePupillaryDistance           = "00460060"
	NearPupillaryDistance               = "00460062"
	OtherPupillaryDistance              = "00460064"
	RadiusOfCurvature                   = "00460075"
	KeratometricPower                   = "00460076"
	KeratometricAxis                    = "00460077"
	BackgroundColor                     = "00460092"
	Optotype                            = "00460094"
	OptotypePresentation                = "00460095"
	AddNearSequence                     = "00460100"
	AddIntermediateSequence             = "00460101"
	AddOtherSequence                    = "00460102"
	AddPower                            = "00460104"
	ViewingDistance                     = "00460106"
	ViewingDistanceType                 = "00460125"
	VisualAcuityModifiers               = "00460135"
	DecimalVisualAcuity                 = "00460137"
	OptotypeDetailedDefinition          = "00460139"
	SpherePower                         = "00460146"
	CylinderPower                       = "00460147"
	CalibrationImage                    = "00500004"
	DeviceSequence                      = "00500010"
	DeviceLength                        = "00500014"
	ContainerComponentWidth             = "00500015"
	DeviceDiameter                      = "00500016"
	DeviceDiameterUnits                 = "00500017"
	DeviceVolume                        = "00500018"
	InterMarkerDistance                 = "00500019"
	ContainerComponentID                = "0050001B"
	DeviceDescription                   = "00500020"
	EnergyWindowVector                  = "00540010"
	NumberOfEnergyWindows               = "00540011"
	EnergyWindowInformationSequence     = "00540012"
	EnergyWindowRangeSequence           = "00540013"
	EnergyWindowLowerLimit              = "00540014"
	EnergyWindowUpperLimit              = "00540015"
	RadiopharmaceuticalInformationSeq   = "00540016"
	ResidualSyringeCounts               = "00540017"
	EnergyWindowName                    = "00540018"
	DetectorVector                      = "00540020"
	NumberOfDetectors                   = "00540021"
	DetectorInformationSequence         = "00540022"
	PhaseVector                         = "00540030"
	NumberOfPhases                      = "00540031"
	PhaseInformationSequence            = "00540032"
	NumberOfFramesInPhase               = "00540033"
	PhaseDelay                          = "00540036"
	PauseBetweenFrames                  = "00540038"
	PhaseDescription                    = "00540039"
	RotationVector                      = "00540050"
	NumberOfRotations                   = "00540051"
	RotationInformationSequence         = "00540052"
	NumberOfFramesInRotation            = "00540053"
	RRIntervalVector                    = "00540060"
	NumberOfRRIntervals                 = "00540061"
	GatedInformationSequence            = "00540062"
	DataInformationSequence             = "00540063"
	TimeSlotVector                      = "00540070"
	NumberOfTimeSlots                   = "00540071"
	TimeSlotInformationSequence         = "00540072"
	TimeSlotTime                        = "00540073"
	SliceVector                         = "00540080"
	NumberOfSlices                      = "00540081"
	AngularViewVector                   = "00540090"
	TimeSliceVector                     = "00540100"
	NumberOfTimeSlices                  = "00540101"
	StartAngle                          = "00540200"
	TypeOfDetectorMotion                = "00540202"
	TriggerVector                       = "00540210"
	NumberOfTriggersInPhase             = "00540211"
	ViewCodeSequence                    = "00540220"
	ViewModifierCodeSequence            = "00540222"
	RadionuclideCodeSequence            = "00540300"
	AdministrationRouteCodeSequence     = "00540302"
	RadiopharmaceuticalCodeSequence     = "00540304"
	CalibrationDataSequence             = "00540306"
	EnergyWindowNumber                  = "00540308"
	ImageID                             = "00540400"
	PatientOrientationCodeSequence      = "00540410"
	PatientOrientationModifierCodeSeq   = "00540412"
	PatientGantryRelationshipCodeSeq    = "00540414"
	SliceProgressionDirection           = "00540500"
	SeriesType                          = "00541000"
	Units                               = "00541001"
	CountsSource                        = "00541002"
	ReprojectionMethod                  = "00541004"
	RandomsCorrectionMethod             = "00541100"
	AttenuationCorrectionMethod         = "00541101"
	DecayCorrection                     = "00541102"
	ReconstructionMethod                = "00541103"
	DetectorLinesOfResponseUsed         = "00541104"
	ScatterCorrectionMethod             = "00541105"
	AxialAcceptance                     = "00541200"
	AxialMash                           = "00541201"
	TransverseMash                      = "00541202"
	DetectorElementSize                 = "00541203"
	CoincidenceWindowWidth              = "00541210"
	SecondaryCountsType                 = "00541220"
	FrameReferenceTime                  = "00541300"
	PrimaryCountsAccumulated            = "00541310"
	SecondaryCountsAccumulated          = "00541311"
	SliceSensitivityFactor              = "00541320"
	DecayFactor                         = "00541321"
	DoseCalibrationFactor               = "00541322"
	ScatterFractionFactor               = "00541323"
	DeadTimeFactor                      = "00541324"
	ImageIndex                          = "00541330"
	CountsIncluded                      = "00541400"
	DeadTimeCorrectionFlag              = "00541401"
	HistogramSequence                   = "00603000"
	HistogramNumberOfBins               = "00603002"
	HistogramFirstBinValue              = "00603004"
	HistogramLastBinValue               = "00603006"
	HistogramBinWidth                   = "00603008"
	HistogramExplanation                = "00603010"
	HistogramData                       = "00603020"
	SegmentationType                    = "00620001"
	SegmentSequence                     = "00620002"
	SegmentedPropertyCategoryCodeSeq    = "00620003"
	SegmentLabel                        = "00620005"
	SegmentDescription                  = "00620006"
	SegmentAlgorithmType                = "00620008"
	SegmentAlgorithmName                = "00620009"
	SegmentIdentificationSequence       = "0062000A"
	ReferencedSegmentNumber             = "0062000B"
	RecommendedDisplayGrayscaleValue    = "0062000C"
	RecommendedDisplayCIELabValue       = "0062000D"
	MaximumFractionalValue              = "0062000E"
	SegmentedPropertyTypeCodeSequence   = "0062000F"
	SegmentationFractionalType          = "00620010"
	DeformableRegistrationSequence      = "00640002"
	SourceFrameOfReferenceUID           = "00640003"
	DeformableRegistrationGridSequence  = "00640005"
	GridDimensions                      = "00640007"
	GridResolution                      = "00640008"
	VectorGridData                      = "00640009"
	PreDeformationMatrixRegistSeq       = "0064000F"
	PostDeformationMatrixRegistSeq      = "00640010"
	NumberOfSurfaces                    = "00660001"
	SurfaceSequence                     = "00660002"
	SurfaceNumber                       = "00660003"
	SurfaceComments                     = "00660004"
	SurfaceProcessing                   = "00660009"
	SurfaceProcessingRatio              = "0066000A"
	FiniteVolume                        = "0066000E"
	Manifold                            = "00660010"
	SurfacePointsSequence               = "00660011"
	NumberOfSurfacePoints               = "00660015"
	PointCoordinatesData                = "00660016"
	PointPositionAccuracy               = "00660017"
	MeanPointDistance                   = "00660018"
	MaximumPointDistance                = "00660019"
	AxisOfRotation                      = "0066001B"
	CenterOfRotation                    = "0066001C"
	NumberOfVectors                     = "0066001E"
	VectorDimensionality                = "0066001F"
	VectorAccuracy                      = "00660020"
	VectorCoordinateData                = "00660021"
	TrianglePointIndexList              = "00660023"
	EdgePointIndexList                  = "00660024"
	VertexPointIndexList                = "00660025"
	TriangleStripSequence               = "00660026"
	TriangleFanSequence                 = "00660027"
	LineSequence                        = "00660028"
	PrimitivePointIndexList             = "00660029"
	SurfaceCount                        = "0066002A"
	AlgorithmFamilyCodeSequ             = "0066002F"
	AlgorithmVersion                    = "00660031"
	AlgorithmParameters                 = "00660032"
	FacetSequence                       = "00660034"
	AlgorithmName                       = "00660036"
	GraphicAnnotationSequence           = "00700001"
	GraphicLayer                        = "00700002"
	BoundingBoxAnnotationUnits          = "00700003"
	AnchorPointAnnotationUnits          = "00700004"
	GraphicAnnotationUnits              = "00700005"
	UnformattedTextValue                = "00700006"
	TextObjectSequence                  = "00700008"
	GraphicObjectSequence               = "00700009"
	BoundingBoxTopLeftHandCorner        = "00700010"
	BoundingBoxBottomRightHandCorner    = "00700011"
	BoundingBoxTextHorizJustification   = "00700012"
	AnchorPoint                         = "00700014"
	AnchorPointVisibility               = "00700015"
	GraphicDimensions                   = "00700020"
	NumberOfGraphicPoints               = "00700021"
	GraphicData                         = "00700022"
	GraphicType                         = "00700023"
	GraphicFilled                       = "00700024"
	ImageRotationRetired                = "00700040"
	ImageHorizontalFlip                 = "00700041"
	ImageRotation                       = "00700042"
	DisplayedAreaTopLeftTrial           = "00700050"
	DisplayedAreaBottomRightTrial       = "00700051"
	DisplayedAreaTopLeft                = "00700052"
	DisplayedAreaBottomRight            = "00700053"
	DisplayedAreaSelectionSequence      = "0070005A"
	GraphicLayerSequence                = "00700060"
	GraphicLayerOrder                   = "00700062"
	GraphicLayerRecDisplayGraysclValue  = "00700066"
	GraphicLayerRecDisplayRGBValue      = "00700067"
	GraphicLayerDescription             = "00700068"
	ContentLabel                        = "00700080"
	ContentDescription                  = "00700081"
	PresentationCreationDate            = "00700082"
	PresentationCreationTime            = "00700083"
	ContentCreatorName                  = "00700084"
	ContentCreatorIDCodeSequence        = "00700086"
	PresentationSizeMode                = "00700100"
	PresentationPixelSpacing            = "00700101"
	PresentationPixelAspectRatio        = "00700102"
	PresentationPixelMagRatio           = "00700103"
	ShapeType                           = "00700306"
	RegistrationSequence                = "00700308"
	MatrixRegistrationSequence          = "00700309"
	MatrixSequence                      = "0070030A"
	FrameOfRefTransformationMatrixType  = "0070030C"
	RegistrationTypeCodeSequence        = "0070030D"
	FiducialDescription                 = "0070030F"
	FiducialIdentifier                  = "00700310"
	FiducialIdentifierCodeSequence      = "00700311"
	ContourUncertaintyRadius            = "00700312"
	UsedFiducialsSequence               = "00700314"
	GraphicCoordinatesDataSequence      = "00700318"
	FiducialUID                         = "0070031A"
	FiducialSetSequence                 = "0070031C"
	FiducialSequence                    = "0070031E"
	GraphicLayerRecomDisplayCIELabVal   = "00700401"
	BlendingSequence                    = "00700402"
	RelativeOpacity                     = "00700403"
	ReferencedSpatialRegistrationSeq    = "00700404"
	BlendingPosition                    = "00700405"
	HangingProtocolName                 = "00720002"
	HangingProtocolDescription          = "00720004"
	HangingProtocolLevel                = "00720006"
	HangingProtocolCreator              = "00720008"
	HangingProtocolCreationDateTime     = "0072000A"
	HangingProtocolDefinitionSequence   = "0072000C"
	HangingProtocolUserIDCodeSequence   = "0072000E"
	HangingProtocolUserGroupName        = "00720010"
	SourceHangingProtocolSequence       = "00720012"
	NumberOfPriorsReferenced            = "00720014"
	ImageSetsSequence                   = "00720020"
	ImageSetSelectorSequence            = "00720022"
	ImageSetSelectorUsageFlag           = "00720024"
	SelectorAttribute                   = "00720026"
	SelectorValueNumber                 = "00720028"
	TimeBasedImageSetsSequence          = "00720030"
	ImageSetNumber                      = "00720032"
	ImageSetSelectorCategory            = "00720034"
	RelativeTime                        = "00720038"
	RelativeTimeUnits                   = "0072003A"
	AbstractPriorValue                  = "0072003C"
	AbstractPriorCodeSequence           = "0072003E"
	ImageSetLabel                       = "00720040"
	SelectorAttributeVR                 = "00720050"
	SelectorSequencePointer             = "00720052"
	SelectorSeqPointerPrivateCreator    = "00720054"
	SelectorAttributePrivateCreator     = "00720056"
	SelectorATValue                     = "00720060"
	SelectorCSValue                     = "00720062"
	SelectorISValue                     = "00720064"
	SelectorLOValue                     = "00720066"
	SelectorLTValue                     = "00720068"
	SelectorPNValue                     = "0072006A"
	SelectorSHValue                     = "0072006C"
	SelectorSTValue                     = "0072006E"
	SelectorUTValue                     = "00720070"
	SelectorDSValue                     = "00720072"
	SelectorFDValue                     = "00720074"
	SelectorFLValue                     = "00720076"
	SelectorULValue                     = "00720078"
	SelectorUSValue                     = "0072007A"
	SelectorSLValue                     = "0072007C"
	SelectorSSValue                     = "0072007E"
	SelectorCodeSequenceValue           = "00720080"
	NumberOfScreens                     = "00720100"
	NominalScreenDefinitionSequence     = "00720102"
	NumberOfVerticalPixels              = "00720104"
	NumberOfHorizontalPixels            = "00720106"
	DisplayEnvironmentSpatialPosition   = "00720108"
	ScreenMinimumGrayscaleBitDepth      = "0072010A"
	ScreenMinimumColorBitDepth          = "0072010C"
	ApplicationMaximumRepaintTime       = "0072010E"
	DisplaySetsSequence                 = "00720200"
	DisplaySetNumber                    = "00720202"
	DisplaySetLabel                     = "00720203"
	DisplaySetPresentationGroup         = "00720204"
	DisplaySetPresentationGroupDescr    = "00720206"
	PartialDataDisplayHandling          = "00720208"
	SynchronizedScrollingSequence       = "00720210"
	DisplaySetScrollingGroup            = "00720212"
	NavigationIndicatorSequence         = "00720214"
	NavigationDisplaySet                = "00720216"
	ReferenceDisplaySets                = "00720218"
	ImageBoxesSequence                  = "00720300"
	ImageBoxNumber                      = "00720302"
	ImageBoxLayoutType                  = "00720304"
	ImageBoxTileHorizontalDimension     = "00720306"
	ImageBoxTileVerticalDimension       = "00720308"
	ImageBoxScrollDirection             = "00720310"
	ImageBoxSmallScrollType             = "00720312"
	ImageBoxSmallScrollAmount           = "00720314"
	ImageBoxLargeScrollType             = "00720316"
	ImageBoxLargeScrollAmount           = "00720318"
	ImageBoxOverlapPriority             = "00720320"
	CineRelativeToRealTime              = "00720330"
	FilterOperationsSequence            = "00720400"
	FilterByCategory                    = "00720402"
	FilterByAttributePresence           = "00720404"
	FilterByOperator                    = "00720406"
	SynchronizedImageBoxList            = "00720432"
	TypeOfSynchronization               = "00720434"
	BlendingOperationType               = "00720500"
	ReformattingOperationType           = "00720510"
	ReformattingThickness               = "00720512"
	ReformattingInterval                = "00720514"
	ReformattingOpInitialViewDir        = "00720516"
	RenderingType3D                     = "00720520"
	SortingOperationsSequence           = "00720600"
	SortByCategory                      = "00720602"
	SortingDirection                    = "00720604"
	DisplaySetPatientOrientation        = "00720700"
	VOIType                             = "00720702"
	PseudoColorType                     = "00720704"
	ShowGrayscaleInverted               = "00720706"
	ShowImageTrueSizeFlag               = "00720710"
	ShowGraphicAnnotationFlag           = "00720712"
	ShowPatientDemographicsFlag         = "00720714"
	ShowAcquisitionTechniquesFlag       = "00720716"
	DisplaySetHorizontalJustification   = "00720717"
	DisplaySetVerticalJustification     = "00720718"
	UnifiedProcedureStepState           = "00741000"
	UPSProgressInformationSequence      = "00741002"
	UnifiedProcedureStepProgress        = "00741004"
	UnifiedProcedureStepProgressDescr   = "00741006"
	UnifiedProcedureStepComURISeq       = "00741008"
	ContactURI                          = "0074100a"
	ContactDisplayName                  = "0074100c"
	BeamTaskSequence                    = "00741020"
	BeamTaskType                        = "00741022"
	BeamOrderIndex                      = "00741024"
	DeliveryVerificationImageSequence   = "00741030"
	VerificationImageTiming             = "00741032"
	DoubleExposureFlag                  = "00741034"
	DoubleExposureOrdering              = "00741036"
	DoubleExposureMeterset              = "00741038"
	DoubleExposureFieldDelta            = "0074103A"
	RelatedReferenceRTImageSequence     = "00741040"
	GeneralMachineVerificationSequence  = "00741042"
	ConventionalMachineVerificationSeq  = "00741044"
	IonMachineVerificationSequence      = "00741046"
	FailedAttributesSequence            = "00741048"
	OverriddenAttributesSequence        = "0074104A"
	ConventionalControlPointVerifySeq   = "0074104C"
	IonControlPointVerificationSeq      = "0074104E"
	AttributeOccurrenceSequence         = "00741050"
	AttributeOccurrencePointer          = "00741052"
	AttributeItemSelector               = "00741054"
	AttributeOccurrencePrivateCreator   = "00741056"
	ScheduledProcedureStepPriority      = "00741200"
	WorklistLabel                       = "00741202"
	ProcedureStepLabel                  = "00741204"
	ScheduledProcessingParametersSeq    = "00741210"
	PerformedProcessingParametersSeq    = "00741212"
	UPSPerformedProcedureSequence       = "00741216"
	RelatedProcedureStepSequence        = "00741220"
	ProcedureStepRelationshipType       = "00741222"
	DeletionLock                        = "00741230"
	ReceivingAE                         = "00741234"
	RequestingAE                        = "00741236"
	ReasonForCancellation               = "00741238"
	SCPStatus                           = "00741242"
	SubscriptionListStatus              = "00741244"
	UPSListStatus                       = "00741246"
	StorageMediaFileSetID               = "00880130"
	StorageMediaFileSetUID              = "00880140"
	IconImageSequence                   = "00880200"
	TopicTitle                          = "00880904"
	TopicSubject                        = "00880906"
	TopicAuthor                         = "00880910"
	TopicKeywords                       = "00880912"
	SOPInstanceStatus                   = "01000410"
	SOPAuthorizationDateAndTime         = "01000420"
	SOPAuthorizationComment             = "01000424"
	AuthorizationEquipmentCertNumber    = "01000426"
	MACIDNumber                         = "04000005"
	MACCalculationTransferSyntaxUID     = "04000010"
	MACAlgorithm                        = "04000015"
	DataElementsSigned                  = "04000020"
	DigitalSignatureUID                 = "04000100"
	DigitalSignatureDateTime            = "04000105"
	CertificateType                     = "04000110"
	CertificateOfSigner                 = "04000115"
	Signature                           = "04000120"
	CertifiedTimestampType              = "04000305"
	CertifiedTimestamp                  = "04000310"
	DigitalSignaturePurposeCodeSeq      = "04000401"
	ReferencedDigitalSignatureSeq       = "04000402"
	ReferencedSOPInstanceMACSeq         = "04000403"
	MAC                                 = "04000404"
	EncryptedAttributesSequence         = "04000500"
	EncryptedContentTransferSyntaxUID   = "04000510"
	EncryptedContent                    = "04000520"
	ModifiedAttributesSequence          = "04000550"
	OriginalAttributesSequence          = "04000561"
	AttributeModificationDateTime       = "04000562"
	ModifyingSystem                     = "04000563"
	SourceOfPreviousValues              = "04000564"
	ReasonForTheAttributeModification   = "04000565"
	EscapeTriplet                       = "1000xxx0"
	RunLengthTriplet                    = "1000xxx1"
	HuffmanTableSize                    = "1000xxx2"
	HuffmanTableTriplet                 = "1000xxx3"
	ShiftTableSize                      = "1000xxx4"
	ShiftTableTriplet                   = "1000xxx5"
	ZonalMap                            = "1010xxxx"
	NumberOfCopies                      = "20000010"
	PrinterConfigurationSequence        = "2000001E"
	PrintPriority                       = "20000020"
	MediumType                          = "20000030"
	FilmDestination                     = "20000040"
	FilmSessionLabel                    = "20000050"
	MemoryAllocation                    = "20000060"
	MaximumMemoryAllocation             = "20000061"
	ColorImagePrintingFlag              = "20000062"
	CollationFlag                       = "20000063"
	AnnotationFlag                      = "20000065"
	ImageOverlayFlag                    = "20000067"
	PresentationLUTFlag                 = "20000069"
	ImageBoxPresentationLUTFlag         = "2000006A"
	MemoryBitDepth                      = "200000A0"
	PrintingBitDepth                    = "200000A1"
	MediaInstalledSequence              = "200000A2"
	OtherMediaAvailableSequence         = "200000A4"
	SupportedImageDisplayFormatSeq      = "200000A8"
	ReferencedFilmBoxSequence           = "20000500"
	ReferencedStoredPrintSequence       = "20000510"
	ImageDisplayFormat                  = "20100010"
	AnnotationDisplayFormatID           = "20100030"
	FilmOrientation                     = "20100040"
	FilmSizeID                          = "20100050"
	PrinterResolutionID                 = "20100052"
	DefaultPrinterResolutionID          = "20100054"
	MagnificationType                   = "20100060"
	SmoothingType                       = "20100080"
	DefaultMagnificationType            = "201000A6"
	OtherMagnificationTypesAvailable    = "201000A7"
	DefaultSmoothingType                = "201000A8"
	OtherSmoothingTypesAvailable        = "201000A9"
	BorderDensity                       = "20100100"
	EmptyImageDensity                   = "20100110"
	MinDensity                          = "20100120"
	MaxDensity                          = "20100130"
	Trim                                = "20100140"
	ConfigurationInformation            = "20100150"
	ConfigurationInformationDescr       = "20100152"
	MaximumCollatedFilms                = "20100154"
	Illumination                        = "2010015E"
	ReflectedAmbientLight               = "20100160"
	PrinterPixelSpacing                 = "20100376"
	ReferencedFilmSessionSequence       = "20100500"
	ReferencedImageBoxSequence          = "20100510"
	ReferencedBasicAnnotationBoxSeq     = "20100520"
	ImageBoxPosition                    = "20200010"
	Polarity                            = "20200020"
	RequestedImageSize                  = "20200030"
	RequestedDecimateCropBehavior       = "20200040"
	RequestedResolutionID               = "20200050"
	RequestedImageSizeFlag              = "202000A0"
	DecimateCropResult                  = "202000A2"
	BasicGrayscaleImageSequence         = "20200110"
	BasicColorImageSequence             = "20200111"
	ReferencedImageOverlayBoxSequence   = "20200130"
	ReferencedVOILUTBoxSequence         = "20200140"
	AnnotationPosition                  = "20300010"
	TextString                          = "20300020"
	ReferencedOverlayPlaneSequence      = "20400010"
	ReferencedOverlayPlaneGroups        = "20400011"
	OverlayPixelDataSequence            = "20400020"
	OverlayMagnificationType            = "20400060"
	OverlaySmoothingType                = "20400070"
	OverlayOrImageMagnification         = "20400072"
	MagnifyToNumberOfColumns            = "20400074"
	OverlayForegroundDensity            = "20400080"
	OverlayBackgroundDensity            = "20400082"
	OverlayMode                         = "20400090"
	ThresholdDensity                    = "20400100"
	PresentationLUTSequence             = "20500010"
	PresentationLUTShape                = "20500020"
	ReferencedPresentationLUTSequence   = "20500500"
	PrintJobID                          = "21000010"
	ExecutionStatus                     = "21000020"
	ExecutionStatusInfo                 = "21000030"
	CreationDate                        = "21000040"
	CreationTime                        = "21000050"
	Originator                          = "21000070"
	DestinationAE                       = "21000140"
	OwnerID                             = "21000160"
	NumberOfFilms                       = "21000170"
	ReferencedPrintJobSequence          = "21000500"
	PrinterStatus                       = "21100010"
	PrinterStatusInfo                   = "21100020"
	PrinterName                         = "21100030"
	PrintQueueID                        = "21100099"
	QueueStatus                         = "21200010"
	PrintJobDescriptionSequence         = "21200050"
	PrintManagementCapabilitiesSeq      = "21300010"
	PrinterCharacteristicsSequence      = "21300015"
	FilmBoxContentSequence              = "21300030"
	ImageBoxContentSequence             = "21300040"
	AnnotationContentSequence           = "21300050"
	ImageOverlayBoxContentSequence      = "21300060"
	PresentationLUTContentSequence      = "21300080"
	ProposedStudySequence               = "213000A0"
	OriginalImageSequence               = "213000C0"
	LabelFromInfoExtractedFromInstance  = "22000001"
	LabelText                           = "22000002"
	LabelStyleSelection                 = "22000003"
	MediaDisposition                    = "22000004"
	BarcodeValue                        = "22000005"
	BarcodeSymbology                    = "22000006"
	AllowMediaSplitting                 = "22000007"
	IncludeNonDICOMObjects              = "22000008"
	IncludeDisplayApplication           = "22000009"
	SaveCompInstancesAfterMediaCreate   = "2200000A"
	TotalNumberMediaPiecesCreated       = "2200000B"
	RequestedMediaApplicationProfile    = "2200000C"
	ReferencedStorageMediaSequence      = "2200000D"
	FailureAttributes                   = "2200000E"
	AllowLossyCompression               = "2200000F"
	RequestPriority                     = "22000020"
	RTImageLabel                        = "30020002"
	RTImageName                         = "30020003"
	RTImageDescription                  = "30020004"
	ReportedValuesOrigin                = "3002000A"
	RTImagePlane                        = "3002000C"
	XRayImageReceptorTranslation        = "3002000D"
	XRayImageReceptorAngle              = "3002000E"
	RTImageOrientation                  = "30020010"
	ImagePlanePixelSpacing              = "30020011"
	RTImagePosition                     = "30020012"
	RadiationMachineName                = "30020020"
	RadiationMachineSAD                 = "30020022"
	RadiationMachineSSD                 = "30020024"
	RTImageSID                          = "30020026"
	SourceToReferenceObjectDistance     = "30020028"
	FractionNumber                      = "30020029"
	ExposureSequence                    = "30020030"
	MetersetExposure                    = "30020032"
	DiaphragmPosition                   = "30020034"
	FluenceMapSequence                  = "30020040"
	FluenceDataSource                   = "30020041"
	FluenceDataScale                    = "30020042"
	FluenceMode                         = "30020051"
	FluenceModeID                       = "30020052"
	DVHType                             = "30040001"
	DoseUnits                           = "30040002"
	DoseType                            = "30040004"
	DoseComment                         = "30040006"
	NormalizationPoint                  = "30040008"
	DoseSummationType                   = "3004000A"
	GridFrameOffsetVector               = "3004000C"
	DoseGridScaling                     = "3004000E"
	RTDoseROISequence                   = "30040010"
	DoseValue                           = "30040012"
	TissueHeterogeneityCorrection       = "30040014"
	DVHNormalizationPoint               = "30040040"
	DVHNormalizationDoseValue           = "30040042"
	DVHSequence                         = "30040050"
	DVHDoseScaling                      = "30040052"
	DVHVolumeUnits                      = "30040054"
	DVHNumberOfBins                     = "30040056"
	DVHData                             = "30040058"
	DVHReferencedROISequence            = "30040060"
	DVHROIContributionType              = "30040062"
	DVHMinimumDose                      = "30040070"
	DVHMaximumDose                      = "30040072"
	DVHMeanDose                         = "30040074"
	StructureSetLabel                   = "30060002"
	StructureSetName                    = "30060004"
	StructureSetDescription             = "30060006"
	StructureSetDate                    = "30060008"
	StructureSetTime                    = "30060009"
	ReferencedFrameOfReferenceSequence  = "30060010"
	RTReferencedStudySequence           = "30060012"
	RTReferencedSeriesSequence          = "30060014"
	ContourImageSequence                = "30060016"
	StructureSetROISequence             = "30060020"
	ROINumber                           = "30060022"
	ReferencedFrameOfReferenceUID       = "30060024"
	ROIName                             = "30060026"
	ROIDescription                      = "30060028"
	ROIDisplayColor                     = "3006002A"
	ROIVolume                           = "3006002C"
	RTRelatedROISequence                = "30060030"
	RTROIRelationship                   = "30060033"
	ROIGenerationAlgorithm              = "30060036"
	ROIGenerationDescription            = "30060038"
	ROIContourSequence                  = "30060039"
	ContourSequence                     = "30060040"
	ContourGeometricType                = "30060042"
	ContourSlabThickness                = "30060044"
	ContourOffsetVector                 = "30060045"
	NumberOfContourPoints               = "30060046"
	ContourNumber                       = "30060048"
	AttachedContours                    = "30060049"
	ContourData                         = "30060050"
	RTROIObservationsSequence           = "30060080"
	ObservationNumber                   = "30060082"
	ReferencedROINumber                 = "30060084"
	ROIObservationLabel                 = "30060085"
	RTROIIdentificationCodeSequence     = "30060086"
	ROIObservationDescription           = "30060088"
	RelatedRTROIObservationsSequence    = "300600A0"
	RTROIInterpretedType                = "300600A4"
	ROIInterpreter                      = "300600A6"
	ROIPhysicalPropertiesSequence       = "300600B0"
	ROIPhysicalProperty                 = "300600B2"
	ROIPhysicalPropertyValue            = "300600B4"
	ROIElementalCompositionSequence     = "300600B6"
	ROIElementalCompAtomicNumber        = "300600B7"
	ROIElementalCompAtomicMassFraction  = "300600B8"
	FrameOfReferenceRelationshipSeq     = "300600C0"
	RelatedFrameOfReferenceUID          = "300600C2"
	FrameOfReferenceTransformType       = "300600C4"
	FrameOfReferenceTransformMatrix     = "300600C6"
	FrameOfReferenceTransformComment    = "300600C8"
	MeasuredDoseReferenceSequence       = "30080010"
	MeasuredDoseDescription             = "30080012"
	MeasuredDoseType                    = "30080014"
	MeasuredDoseValue                   = "30080016"
	TreatmentSessionBeamSequence        = "30080020"
	TreatmentSessionIonBeamSequence     = "30080021"
	CurrentFractionNumber               = "30080022"
	TreatmentControlPointDate           = "30080024"
	TreatmentControlPointTime           = "30080025"
	TreatmentTerminationStatus          = "3008002A"
	TreatmentTerminationCode            = "3008002B"
	TreatmentVerificationStatus         = "3008002C"
	ReferencedTreatmentRecordSequence   = "30080030"
	SpecifiedPrimaryMeterset            = "30080032"
	SpecifiedSecondaryMeterset          = "30080033"
	DeliveredPrimaryMeterset            = "30080036"
	DeliveredSecondaryMeterset          = "30080037"
	SpecifiedTreatmentTime              = "3008003A"
	DeliveredTreatmentTime              = "3008003B"
	ControlPointDeliverySequence        = "30080040"
	IonControlPointDeliverySequence     = "30080041"
	SpecifiedMeterset                   = "30080042"
	DeliveredMeterset                   = "30080044"
	MetersetRateSet                     = "30080045"
	MetersetRateDelivered               = "30080046"
	ScanSpotMetersetsDelivered          = "30080047"
	DoseRateDelivered                   = "30080048"
	TreatmentSummaryCalcDoseRefSeq      = "30080050"
	CumulativeDoseToDoseReference       = "30080052"
	FirstTreatmentDate                  = "30080054"
	MostRecentTreatmentDate             = "30080056"
	NumberOfFractionsDelivered          = "3008005A"
	OverrideSequence                    = "30080060"
	ParameterSequencePointer            = "30080061"
	OverrideParameterPointer            = "30080062"
	ParameterItemIndex                  = "30080063"
	MeasuredDoseReferenceNumber         = "30080064"
	ParameterPointer                    = "30080065"
	OverrideReason                      = "30080066"
	CorrectedParameterSequence          = "30080068"
	CorrectionValue                     = "3008006A"
	CalculatedDoseReferenceSequence     = "30080070"
	CalculatedDoseReferenceNumber       = "30080072"
	CalculatedDoseReferenceDescription  = "30080074"
	CalculatedDoseReferenceDoseValue    = "30080076"
	StartMeterset                       = "30080078"
	EndMeterset                         = "3008007A"
	ReferencedMeasuredDoseReferenceSeq  = "30080080"
	ReferencedMeasuredDoseReferenceNum  = "30080082"
	ReferencedCalculatedDoseRefSeq      = "30080090"
	ReferencedCalculatedDoseRefNumber   = "30080092"
	BeamLimitingDeviceLeafPairsSeq      = "300800A0"
	RecordedWedgeSequence               = "300800B0"
	RecordedCompensatorSequence         = "300800C0"
	RecordedBlockSequence               = "300800D0"
	TreatmentSummaryMeasuredDoseRefSeq  = "300800E0"
	RecordedSnoutSequence               = "300800F0"
	RecordedRangeShifterSequence        = "300800F2"
	RecordedLateralSpreadingDeviceSeq   = "300800F4"
	RecordedRangeModulatorSequence      = "300800F6"
	RecordedSourceSequence              = "30080100"
	SourceSerialNumber                  = "30080105"
	TreatmentSessionAppSetupSeq         = "30080110"
	ApplicationSetupCheck               = "30080116"
	RecordedBrachyAccessoryDeviceSeq    = "30080120"
	ReferencedBrachyAccessoryDeviceNum  = "30080122"
	RecordedChannelSequence             = "30080130"
	SpecifiedChannelTotalTime           = "30080132"
	DeliveredChannelTotalTime           = "30080134"
	SpecifiedNumberOfPulses             = "30080136"
	DeliveredNumberOfPulses             = "30080138"
	SpecifiedPulseRepetitionInterval    = "3008013A"
	DeliveredPulseRepetitionInterval    = "3008013C"
	RecordedSourceApplicatorSequence    = "30080140"
	ReferencedSourceApplicatorNumber    = "30080142"
	RecordedChannelShieldSequence       = "30080150"
	ReferencedChannelShieldNumber       = "30080152"
	BrachyControlPointDeliveredSeq      = "30080160"
	SafePositionExitDate                = "30080162"
	SafePositionExitTime                = "30080164"
	SafePositionReturnDate              = "30080166"
	SafePositionReturnTime              = "30080168"
	CurrentTreatmentStatus              = "30080200"
	TreatmentStatusComment              = "30080202"
	FractionGroupSummarySequence        = "30080220"
	ReferencedFractionNumber            = "30080223"
	FractionGroupType                   = "30080224"
	BeamStopperPosition                 = "30080230"
	FractionStatusSummarySequence       = "30080240"
	TreatmentDate                       = "30080250"
	TreatmentTime                       = "30080251"
	RTPlanLabel                         = "300A0002"
	RTPlanName                          = "300A0003"
	RTPlanDescription                   = "300A0004"
	RTPlanDate                          = "300A0006"
	RTPlanTime                          = "300A0007"
	TreatmentProtocols                  = "300A0009"
	PlanIntent                          = "300A000A"
	TreatmentSites                      = "300A000B"
	RTPlanGeometry                      = "300A000C"
	PrescriptionDescription             = "300A000E"
	DoseReferenceSequence               = "300A0010"
	DoseReferenceNumber                 = "300A0012"
	DoseReferenceUID                    = "300A0013"
	DoseReferenceStructureType          = "300A0014"
	NominalBeamEnergyUnit               = "300A0015"
	DoseReferenceDescription            = "300A0016"
	DoseReferencePointCoordinates       = "300A0018"
	NominalPriorDose                    = "300A001A"
	DoseReferenceType                   = "300A0020"
	ConstraintWeight                    = "300A0021"
	DeliveryWarningDose                 = "300A0022"
	DeliveryMaximumDose                 = "300A0023"
	TargetMinimumDose                   = "300A0025"
	TargetPrescriptionDose              = "300A0026"
	TargetMaximumDose                   = "300A0027"
	TargetUnderdoseVolumeFraction       = "300A0028"
	OrganAtRiskFullVolumeDose           = "300A002A"
	OrganAtRiskLimitDose                = "300A002B"
	OrganAtRiskMaximumDose              = "300A002C"
	OrganAtRiskOverdoseVolumeFraction   = "300A002D"
	ToleranceTableSequence              = "300A0040"
	ToleranceTableNumber                = "300A0042"
	ToleranceTableLabel                 = "300A0043"
	GantryAngleTolerance                = "300A0044"
	BeamLimitingDeviceAngleTolerance    = "300A0046"
	BeamLimitingDeviceToleranceSeq      = "300A0048"
	BeamLimitingDevicePositionTol       = "300A004A"
	SnoutPositionTolerance              = "300A004B"
	PatientSupportAngleTolerance        = "300A004C"
	TableTopEccentricAngleTolerance     = "300A004E"
	TableTopPitchAngleTolerance         = "300A004F"
	TableTopRollAngleTolerance          = "300A0050"
	TableTopVerticalPositionTolerance   = "300A0051"
	TableTopLongitudinalPositionTol     = "300A0052"
	TableTopLateralPositionTolerance    = "300A0053"
	RTPlanRelationship                  = "300A0055"
	FractionGroupSequence               = "300A0070"
	FractionGroupNumber                 = "300A0071"
	FractionGroupDescription            = "300A0072"
	NumberOfFractionsPlanned            = "300A0078"
	NumberFractionPatternDigitsPerDay   = "300A0079"
	RepeatFractionCycleLength           = "300A007A"
	FractionPattern                     = "300A007B"
	NumberOfBeams                       = "300A0080"
	BeamDoseSpecificationPoint          = "300A0082"
	BeamDose                            = "300A0084"
	BeamMeterset                        = "300A0086"
	BeamDosePointDepth                  = "300A0088"
	BeamDosePointEquivalentDepth        = "300A0089"
	BeamDosePointSSD                    = "300A008A"
	NumberOfBrachyApplicationSetups     = "300A00A0"
	BrachyAppSetupDoseSpecPoint         = "300A00A2"
	BrachyApplicationSetupDose          = "300A00A4"
	BeamSequence                        = "300A00B0"
	TreatmentMachineName                = "300A00B2"
	PrimaryDosimeterUnit                = "300A00B3"
	SourceAxisDistance                  = "300A00B4"
	BeamLimitingDeviceSequence          = "300A00B6"
	RTBeamLimitingDeviceType            = "300A00B8"
	SourceToBeamLimitingDeviceDistance  = "300A00BA"
	IsocenterToBeamLimitingDeviceDist   = "300A00BB"
	NumberOfLeafJawPairs                = "300A00BC"
	LeafPositionBoundaries              = "300A00BE"
	BeamNumber                          = "300A00C0"
	BeamName                            = "300A00C2"
	BeamDescription                     = "300A00C3"
	BeamType                            = "300A00C4"
	RadiationType                       = "300A00C6"
	HighDoseTechniqueType               = "300A00C7"
	ReferenceImageNumber                = "300A00C8"
	PlannedVerificationImageSequence    = "300A00CA"
	ImagingDeviceSpecificAcqParams      = "300A00CC"
	TreatmentDeliveryType               = "300A00CE"
	NumberOfWedges                      = "300A00D0"
	WedgeSequence                       = "300A00D1"
	WedgeNumber                         = "300A00D2"
	WedgeType                           = "300A00D3"
	WedgeID                             = "300A00D4"
	WedgeAngle                          = "300A00D5"
	WedgeFactor                         = "300A00D6"
	TotalWedgeTrayWaterEquivThickness   = "300A00D7"
	WedgeOrientation                    = "300A00D8"
	IsocenterToWedgeTrayDistance        = "300A00D9"
	SourceToWedgeTrayDistance           = "300A00DA"
	WedgeThinEdgePosition               = "300A00DB"
	BolusID                             = "300A00DC"
	BolusDescription                    = "300A00DD"
	NumberOfCompensators                = "300A00E0"
	MaterialID                          = "300A00E1"
	TotalCompensatorTrayFactor          = "300A00E2"
	CompensatorSequence                 = "300A00E3"
	CompensatorNumber                   = "300A00E4"
	CompensatorID                       = "300A00E5"
	SourceToCompensatorTrayDistance     = "300A00E6"
	CompensatorRows                     = "300A00E7"
	CompensatorColumns                  = "300A00E8"
	CompensatorPixelSpacing             = "300A00E9"
	CompensatorPosition                 = "300A00EA"
	CompensatorTransmissionData         = "300A00EB"
	CompensatorThicknessData            = "300A00EC"
	NumberOfBoli                        = "300A00ED"
	CompensatorType                     = "300A00EE"
	NumberOfBlocks                      = "300A00F0"
	TotalBlockTrayFactor                = "300A00F2"
	TotalBlockTrayWaterEquivThickness   = "300A00F3"
	BlockSequence                       = "300A00F4"
	BlockTrayID                         = "300A00F5"
	SourceToBlockTrayDistance           = "300A00F6"
	IsocenterToBlockTrayDistance        = "300A00F7"
	BlockType                           = "300A00F8"
	AccessoryCode                       = "300A00F9"
	BlockDivergence                     = "300A00FA"
	BlockMountingPosition               = "300A00FB"
	BlockNumber                         = "300A00FC"
	BlockName                           = "300A00FE"
	BlockThickness                      = "300A0100"
	BlockTransmission                   = "300A0102"
	BlockNumberOfPoints                 = "300A0104"
	BlockData                           = "300A0106"
	ApplicatorSequence                  = "300A0107"
	ApplicatorID                        = "300A0108"
	ApplicatorType                      = "300A0109"
	ApplicatorDescription               = "300A010A"
	CumulativeDoseReferenceCoefficient  = "300A010C"
	FinalCumulativeMetersetWeight       = "300A010E"
	NumberOfControlPoints               = "300A0110"
	ControlPointSequence                = "300A0111"
	ControlPointIndex                   = "300A0112"
	NominalBeamEnergy                   = "300A0114"
	DoseRateSet                         = "300A0115"
	WedgePositionSequence               = "300A0116"
	WedgePosition                       = "300A0118"
	BeamLimitingDevicePositionSequence  = "300A011A"
	LeafJawPositions                    = "300A011C"
	GantryAngle                         = "300A011E"
	GantryRotationDirection             = "300A011F"
	BeamLimitingDeviceAngle             = "300A0120"
	BeamLimitingDeviceRotateDirection   = "300A0121"
	PatientSupportAngle                 = "300A0122"
	PatientSupportRotationDirection     = "300A0123"
	TableTopEccentricAxisDistance       = "300A0124"
	TableTopEccentricAngle              = "300A0125"
	TableTopEccentricRotateDirection    = "300A0126"
	TableTopVerticalPosition            = "300A0128"
	TableTopLongitudinalPosition        = "300A0129"
	TableTopLateralPosition             = "300A012A"
	IsocenterPosition                   = "300A012C"
	SurfaceEntryPoint                   = "300A012E"
	SourceToSurfaceDistance             = "300A0130"
	CumulativeMetersetWeight            = "300A0134"
	TableTopPitchAngle                  = "300A0140"
	TableTopPitchRotationDirection      = "300A0142"
	TableTopRollAngle                   = "300A0144"
	TableTopRollRotationDirection       = "300A0146"
	HeadFixationAngle                   = "300A0148"
	GantryPitchAngle                    = "300A014A"
	GantryPitchRotationDirection        = "300A014C"
	GantryPitchAngleTolerance           = "300A014E"
	PatientSetupSequence                = "300A0180"
	PatientSetupNumber                  = "300A0182"
	PatientSetupLabel                   = "300A0183"
	PatientAdditionalPosition           = "300A0184"
	FixationDeviceSequence              = "300A0190"
	FixationDeviceType                  = "300A0192"
	FixationDeviceLabel                 = "300A0194"
	FixationDeviceDescription           = "300A0196"
	FixationDevicePosition              = "300A0198"
	FixationDevicePitchAngle            = "300A0199"
	FixationDeviceRollAngle             = "300A019A"
	ShieldingDeviceSequence             = "300A01A0"
	ShieldingDeviceType                 = "300A01A2"
	ShieldingDeviceLabel                = "300A01A4"
	ShieldingDeviceDescription          = "300A01A6"
	ShieldingDevicePosition             = "300A01A8"
	SetupTechnique                      = "300A01B0"
	SetupTechniqueDescription           = "300A01B2"
	SetupDeviceSequence                 = "300A01B4"
	SetupDeviceType                     = "300A01B6"
	SetupDeviceLabel                    = "300A01B8"
	SetupDeviceDescription              = "300A01BA"
	SetupDeviceParameter                = "300A01BC"
	SetupReferenceDescription           = "300A01D0"
	TableTopVerticalSetupDisplacement   = "300A01D2"
	TableTopLongitudinalSetupDisplace   = "300A01D4"
	TableTopLateralSetupDisplacement    = "300A01D6"
	BrachyTreatmentTechnique            = "300A0200"
	BrachyTreatmentType                 = "300A0202"
	TreatmentMachineSequence            = "300A0206"
	SourceSequence                      = "300A0210"
	SourceNumber                        = "300A0212"
	SourceType                          = "300A0214"
	SourceManufacturer                  = "300A0216"
	ActiveSourceDiameter                = "300A0218"
	ActiveSourceLength                  = "300A021A"
	SourceEncapsulationNomThickness     = "300A0222"
	SourceEncapsulationNomTransmission  = "300A0224"
	SourceIsotopeName                   = "300A0226"
	SourceIsotopeHalfLife               = "300A0228"
	SourceStrengthUnits                 = "300A0229"
	ReferenceAirKermaRate               = "300A022A"
	SourceStrength                      = "300A022B"
	SourceStrengthReferenceDate         = "300A022C"
	SourceStrengthReferenceTime         = "300A022E"
	ApplicationSetupSequence            = "300A0230"
	ApplicationSetupType                = "300A0232"
	ApplicationSetupNumber              = "300A0234"
	ApplicationSetupName                = "300A0236"
	ApplicationSetupManufacturer        = "300A0238"
	TemplateNumber                      = "300A0240"
	TemplateType                        = "300A0242"
	TemplateName                        = "300A0244"
	TotalReferenceAirKerma              = "300A0250"
	BrachyAccessoryDeviceSequence       = "300A0260"
	BrachyAccessoryDeviceNumber         = "300A0262"
	BrachyAccessoryDeviceID             = "300A0263"
	BrachyAccessoryDeviceType           = "300A0264"
	BrachyAccessoryDeviceName           = "300A0266"
	BrachyAccessoryDeviceNomThickness   = "300A026A"
	BrachyAccessoryDevNomTransmission   = "300A026C"
	ChannelSequence                     = "300A0280"
	ChannelNumber                       = "300A0282"
	ChannelLength                       = "300A0284"
	ChannelTotalTime                    = "300A0286"
	SourceMovementType                  = "300A0288"
	NumberOfPulses                      = "300A028A"
	PulseRepetitionInterval             = "300A028C"
	SourceApplicatorNumber              = "300A0290"
	SourceApplicatorID                  = "300A0291"
	SourceApplicatorType                = "300A0292"
	SourceApplicatorName                = "300A0294"
	SourceApplicatorLength              = "300A0296"
	SourceApplicatorManufacturer        = "300A0298"
	SourceApplicatorWallNomThickness    = "300A029C"
	SourceApplicatorWallNomTrans        = "300A029E"
	SourceApplicatorStepSize            = "300A02A0"
	TransferTubeNumber                  = "300A02A2"
	TransferTubeLength                  = "300A02A4"
	ChannelShieldSequence               = "300A02B0"
	ChannelShieldNumber                 = "300A02B2"
	ChannelShieldID                     = "300A02B3"
	ChannelShieldName                   = "300A02B4"
	ChannelShieldNominalThickness       = "300A02B8"
	ChannelShieldNominalTransmission    = "300A02BA"
	FinalCumulativeTimeWeight           = "300A02C8"
	BrachyControlPointSequence          = "300A02D0"
	ControlPointRelativePosition        = "300A02D2"
	ControlPoint3DPosition              = "300A02D4"
	CumulativeTimeWeight                = "300A02D6"
	CompensatorDivergence               = "300A02E0"
	CompensatorMountingPosition         = "300A02E1"
	SourceToCompensatorDistance         = "300A02E2"
	TotalCompTrayWaterEquivThickness    = "300A02E3"
	IsocenterToCompensatorTrayDistance  = "300A02E4"
	CompensatorColumnOffset             = "300A02E5"
	IsocenterToCompensatorDistances     = "300A02E6"
	CompensatorRelStoppingPowerRatio    = "300A02E7"
	CompensatorMillingToolDiameter      = "300A02E8"
	IonRangeCompensatorSequence         = "300A02EA"
	CompensatorDescription              = "300A02EB"
	RadiationMassNumber                 = "300A0302"
	RadiationAtomicNumber               = "300A0304"
	RadiationChargeState                = "300A0306"
	ScanMode                            = "300A0308"
	VirtualSourceAxisDistances          = "300A030A"
	SnoutSequence                       = "300A030C"
	SnoutPosition                       = "300A030D"
	SnoutID                             = "300A030F"
	NumberOfRangeShifters               = "300A0312"
	RangeShifterSequence                = "300A0314"
	RangeShifterNumber                  = "300A0316"
	RangeShifterID                      = "300A0318"
	RangeShifterType                    = "300A0320"
	RangeShifterDescription             = "300A0322"
	NumberOfLateralSpreadingDevices     = "300A0330"
	LateralSpreadingDeviceSequence      = "300A0332"
	LateralSpreadingDeviceNumber        = "300A0334"
	LateralSpreadingDeviceID            = "300A0336"
	LateralSpreadingDeviceType          = "300A0338"
	LateralSpreadingDeviceDescription   = "300A033A"
	LateralSpreadingDevWaterEquivThick  = "300A033C"
	NumberOfRangeModulators             = "300A0340"
	RangeModulatorSequence              = "300A0342"
	RangeModulatorNumber                = "300A0344"
	RangeModulatorID                    = "300A0346"
	RangeModulatorType                  = "300A0348"
	RangeModulatorDescription           = "300A034A"
	BeamCurrentModulationID             = "300A034C"
	PatientSupportType                  = "300A0350"
	PatientSupportID                    = "300A0352"
	PatientSupportAccessoryCode         = "300A0354"
	FixationLightAzimuthalAngle         = "300A0356"
	FixationLightPolarAngle             = "300A0358"
	MetersetRate                        = "300A035A"
	RangeShifterSettingsSequence        = "300A0360"
	RangeShifterSetting                 = "300A0362"
	IsocenterToRangeShifterDistance     = "300A0364"
	RangeShifterWaterEquivThickness     = "300A0366"
	LateralSpreadingDeviceSettingsSeq   = "300A0370"
	LateralSpreadingDeviceSetting       = "300A0372"
	IsocenterToLateralSpreadingDevDist  = "300A0374"
	RangeModulatorSettingsSequence      = "300A0380"
	RangeModulatorGatingStartValue      = "300A0382"
	RangeModulatorGatingStopValue       = "300A0384"
	IsocenterToRangeModulatorDistance   = "300A038A"
	ScanSpotTuneID                      = "300A0390"
	NumberOfScanSpotPositions           = "300A0392"
	ScanSpotPositionMap                 = "300A0394"
	ScanSpotMetersetWeights             = "300A0396"
	ScanningSpotSize                    = "300A0398"
	NumberOfPaintings                   = "300A039A"
	IonToleranceTableSequence           = "300A03A0"
	IonBeamSequence                     = "300A03A2"
	IonBeamLimitingDeviceSequence       = "300A03A4"
	IonBlockSequence                    = "300A03A6"
	IonControlPointSequence             = "300A03A8"
	IonWedgeSequence                    = "300A03AA"
	IonWedgePositionSequence            = "300A03AC"
	ReferencedSetupImageSequence        = "300A0401"
	SetupImageComment                   = "300A0402"
	MotionSynchronizationSequence       = "300A0410"
	ControlPointOrientation             = "300A0412"
	GeneralAccessorySequence            = "300A0420"
	GeneralAccessoryID                  = "300A0421"
	GeneralAccessoryDescription         = "300A0422"
	GeneralAccessoryType                = "300A0423"
	GeneralAccessoryNumber              = "300A0424"
	ReferencedRTPlanSequence            = "300C0002"
	ReferencedBeamSequence              = "300C0004"
	ReferencedBeamNumber                = "300C0006"
	ReferencedReferenceImageNumber      = "300C0007"
	StartCumulativeMetersetWeight       = "300C0008"
	EndCumulativeMetersetWeight         = "300C0009"
	ReferencedBrachyAppSetupSeq         = "300C000A"
	ReferencedBrachyAppSetupNumber      = "300C000C"
	ReferencedSourceNumber              = "300C000E"
	ReferencedFractionGroupSequence     = "300C0020"
	ReferencedFractionGroupNumber       = "300C0022"
	ReferencedVerificationImageSeq      = "300C0040"
	ReferencedReferenceImageSequence    = "300C0042"
	ReferencedDoseReferenceSequence     = "300C0050"
	ReferencedDoseReferenceNumber       = "300C0051"
	BrachyReferencedDoseReferenceSeq    = "300C0055"
	ReferencedStructureSetSequence      = "300C0060"
	ReferencedPatientSetupNumber        = "300C006A"
	ReferencedDoseSequence              = "300C0080"
	ReferencedToleranceTableNumber      = "300C00A0"
	ReferencedBolusSequence             = "300C00B0"
	ReferencedWedgeNumber               = "300C00C0"
	ReferencedCompensatorNumber         = "300C00D0"
	ReferencedBlockNumber               = "300C00E0"
	ReferencedControlPointIndex         = "300C00F0"
	ReferencedControlPointSequence      = "300C00F2"
	ReferencedStartControlPointIndex    = "300C00F4"
	ReferencedStopControlPointIndex     = "300C00F6"
	ReferencedRangeShifterNumber        = "300C0100"
	ReferencedLateralSpreadingDevNum    = "300C0102"
	ReferencedRangeModulatorNumber      = "300C0104"
	ApprovalStatus                      = "300E0002"
	ReviewDate                          = "300E0004"
	ReviewTime                          = "300E0005"
	ReviewerName                        = "300E0008"
	TextGroupLength                     = "40000000"
	Arbitrary                           = "40000010"
	TextComments                        = "40004000"
	ResultsID                           = "40080040"
	ResultsIDIssuer                     = "40080042"
	ReferencedInterpretationSequence    = "40080050"
	InterpretationRecordedDate          = "40080100"
	InterpretationRecordedTime          = "40080101"
	InterpretationRecorder              = "40080102"
	ReferenceToRecordedSound            = "40080103"
	InterpretationTranscriptionDate     = "40080108"
	InterpretationTranscriptionTime     = "40080109"
	InterpretationTranscriber           = "4008010A"
	InterpretationText                  = "4008010B"
	InterpretationAuthor                = "4008010C"
	InterpretationApproverSequence      = "40080111"
	InterpretationApprovalDate          = "40080112"
	InterpretationApprovalTime          = "40080113"
	PhysicianApprovingInterpretation    = "40080114"
	InterpretationDiagnosisDescription  = "40080115"
	InterpretationDiagnosisCodeSeq      = "40080117"
	ResultsDistributionListSequence     = "40080118"
	DistributionName                    = "40080119"
	DistributionAddress                 = "4008011A"
	InterpretationID                    = "40080200"
	InterpretationIDIssuer              = "40080202"
	InterpretationTypeID                = "40080210"
	InterpretationStatusID              = "40080212"
	Impressions                         = "40080300"
	ResultsComments                     = "40084000"
	MACParametersSequence               = "4FFE0001"
	CurveDimensions                     = "50xx0005"
	NumberOfPoints                      = "50xx0010"
	TypeOfData                          = "50xx0020"
	CurveDescription                    = "50xx0022"
	AxisUnits                           = "50xx0030"
	AxisLabels                          = "50xx0040"
	DataValueRepresentation             = "50xx0103"
	MinimumCoordinateValue              = "50xx0104"
	MaximumCoordinateValue              = "50xx0105"
	CurveRange                          = "50xx0106"
	CurveDataDescriptor                 = "50xx0110"
	CoordinateStartValue                = "50xx0112"
	CoordinateStepValue                 = "50xx0114"
	CurveActivationLayer                = "50xx1001"
	AudioType                           = "50xx2000"
	AudioSampleFormat                   = "50xx2002"
	NumberOfSamples                     = "50xx2006"
	SampleRate                          = "50xx2008"
	TotalTime                           = "50xx200A"
	AudioSampleData                     = "50xx200C"
	AudioComments                       = "50xx200E"
	CurveLabel                          = "50xx2500"
	ReferencedOverlayGroup              = "50xx2610"
	CurveData                           = "50xx3000"
	SharedFunctionalGroupsSequence      = "52009229"
	PerFrameFunctionalGroupsSequence    = "52009230"
	WaveformSequence                    = "54000100"
	ChannelMinimumValue                 = "54000110"
	ChannelMaximumValue                 = "54000112"
	WaveformBitsAllocated               = "54001004"
	WaveformSampleInterpretation        = "54001006"
	WaveformPaddingValue                = "5400100A"
	WaveformData                        = "54001010"
	FirstOrderPhaseCorrectionAngle      = "56000010"
	SpectroscopyData                    = "56000020"
	OverlayGroupLength                  = "60000000"
	OverlayRows                         = "60xx0010"
	OverlayColumns                      = "60xx0011"
	OverlayPlanes                       = "60xx0012"
	NumberOfFramesInOverlay             = "60xx0015"
	OverlayDescription                  = "60xx0022"
	OverlayType                         = "60xx0040"
	OverlaySubtype                      = "60xx0045"
	OverlayOrigin                       = "60xx0050"
	ImageFrameOrigin                    = "60xx0051"
	OverlayPlaneOrigin                  = "60xx0052"
	OverlayCompressionCode              = "60xx0060"
	OverlayCompressionOriginator        = "60xx0061"
	OverlayCompressionLabel             = "60xx0062"
	OverlayCompressionDescription       = "60xx0063"
	OverlayCompressionStepPointers      = "60xx0066"
	OverlayRepeatInterval               = "60xx0068"
	OverlayBitsGrouped                  = "60xx0069"
	OverlayBitsAllocated                = "60xx0100"
	OverlayBitPosition                  = "60xx0102"
	OverlayFormat                       = "60xx0110"
	OverlayLocation                     = "60xx0200"
	OverlayCodeLabel                    = "60xx0800"
	OverlayNumberOfTables               = "60xx0802"
	OverlayCodeTableLocation            = "60xx0803"
	OverlayBitsForCodeWord              = "60xx0804"
	OverlayActivationLayer              = "60xx1001"
	OverlayDescriptorGray               = "60xx1100"
	OverlayDescriptorRed                = "60xx1101"
	OverlayDescriptorGreen              = "60xx1102"
	OverlayDescriptorBlue               = "60xx1103"
	OverlaysGray                        = "60xx1200"
	OverlaysRed                         = "60xx1201"
	OverlaysGreen                       = "60xx1202"
	OverlaysBlue                        = "60xx1203"
	ROIArea                             = "60xx1301"
	ROIMean                             = "60xx1302"
	ROIStandardDeviation                = "60xx1303"
	OverlayLabel                        = "60xx1500"
	OverlayData                         = "60xx3000"
	OverlayComments                     = "60xx4000"
	PixelDataGroupLength                = "7Fxx0000"
	PixelData                           = "7Fxx0010"
	VariableNextDataGroup               = "7Fxx0011"
	VariableCoefficientsSDVN            = "7Fxx0020"
	VariableCoefficientsSDHN            = "7Fxx0030"
	VariableCoefficientsSDDN            = "7Fxx0040"
	DigitalSignaturesSequence           = "FFFAFFFA"
	DataSetTrailingPadding              = "FFFCFFFC"
	StartOfItem                         = "FFFEE000"
	EndOfItems                          = "FFFEE00D"
	EndOfSequence                       = "FFFEE0DD"
)

var TagToName map[string]string

func init() {
	TagToName = make(map[string]string, len(TagNames))
	for name, tag := range TagNames {
		TagToName[tag] = name
	}
}

var TagNames = map[string]string{
	"FileMetaInfoGroupLength":             "00020000",
	"FileMetaInfoVersion":                 "00020001",
	"MediaStorageSOPClassUID":             "00020002",
	"MediaStorageSOPInstanceUID":          "00020003",
	"TransferSyntaxUID":                   "00020010",
	"ImplementationClassUID":              "00020012",
	"ImplementationVersionName":           "00020013",
	"SourceApplicationEntityTitle":        "00020016",
	"PrivateInformationCreatorUID":        "00020100",
	"PrivateInformation":                  "00020102",
	"FileSetID":                           "00041130",
	"FileSetDescriptorFileID":             "00041141",
	"SpecificCharacterSetOfFile":          "00041142",
	"FirstDirectoryRecordOffset":          "00041200",
	"LastDirectoryRecordOffset":           "00041202",
	"FileSetConsistencyFlag":              "00041212",
	"DirectoryRecordSequence":             "00041220",
	"OffsetOfNextDirectoryRecord":         "00041400",
	"RecordInUseFlag":                     "00041410",
	"LowerLevelDirectoryEntityOffset":     "00041420",
	"DirectoryRecordType":                 "00041430",
	"PrivateRecordUID":                    "00041432",
	"ReferencedFileID":                    "00041500",
	"MRDRDirectoryRecordOffset":           "00041504",
	"ReferencedSOPClassUIDInFile":         "00041510",
	"ReferencedSOPInstanceUIDInFile":      "00041511",
	"ReferencedTransferSyntaxUIDInFile":   "00041512",
	"ReferencedRelatedSOPClassUIDInFile":  "0004151A",
	"NumberOfReferences":                  "00041600",
	"IdentifyingGroupLength":              "00080000",
	"LengthToEnd":                         "00080001",
	"SpecificCharacterSet":                "00080005",
	"LanguageCodeSequence":                "00080006",
	"ImageType":                           "00080008",
	"RecognitionCode":                     "00080010",
	"InstanceCreationDate":                "00080012",
	"InstanceCreationTime":                "00080013",
	"InstanceCreatorUID":                  "00080014",
	"SOPClassUID":                         "00080016",
	"SOPInstanceUID":                      "00080018",
	"RelatedGeneralSOPClassUID":           "0008001A",
	"OriginalSpecializedSOPClassUID":      "0008001B",
	"StudyDate":                           "00080020",
	"SeriesDate":                          "00080021",
	"AcquisitionDate":                     "00080022",
	"ContentDate":                         "00080023",
	"OverlayDate":                         "00080024",
	"CurveDate":                           "00080025",
	"AcquisitionDateTime":                 "0008002A",
	"StudyTime":                           "00080030",
	"SeriesTime":                          "00080031",
	"AcquisitionTime":                     "00080032",
	"ContentTime":                         "00080033",
	"OverlayTime":                         "00080034",
	"CurveTime":                           "00080035",
	"DataSetType":                         "00080040",
	"DataSetSubtype":                      "00080041",
	"NuclearMedicineSeriesType":           "00080042",
	"AccessionNumber":                     "00080050",
	"QueryRetrieveLevel":                  "00080052",
	"RetrieveAETitle":                     "00080054",
	"InstanceAvailability":                "00080056",
	"FailedSOPInstanceUIDList":            "00080058",
	"Modality":                            "00080060",
	"ModalitiesInStudy":                   "00080061",
	"SOPClassesInStudy":                   "00080062",
	"ConversionType":                      "00080064",
	"PresentationIntentType":              "00080068",
	"Manufacturer":                        "00080070",
	"InstitutionName":                     "00080080",
	"InstitutionAddress":                  "00080081",
	"InstitutionCodeSequence":             "00080082",
	"ReferringPhysicianName":              "00080090",
	"ReferringPhysicianAddress":           "00080092",
	"ReferringPhysicianTelephoneNumber":   "00080094",
	"ReferringPhysicianIDSequence":        "00080096",
	"CodeValue":                           "00080100",
	"CodingSchemeDesignator":              "00080102",
	"CodingSchemeVersion":                 "00080103",
	"CodeMeaning":                         "00080104",
	"MappingResource":                     "00080105",
	"ContextGroupVersion":                 "00080106",
	"ContextGroupLocalVersion":            "00080107",
	"ContextGroupExtensionFlag":           "0008010B",
	"CodingSchemeUID":                     "0008010C",
	"ContextGroupExtensionCreatorUID":     "0008010D",
	"ContextIdentifier":                   "0008010F",
	"CodingSchemeIDSequence":              "00080110",
	"CodingSchemeRegistry":                "00080112",
	"CodingSchemeExternalID":              "00080114",
	"CodingSchemeName":                    "00080115",
	"CodingSchemeResponsibleOrganization": "00080116",
	"ContextUID":                          "00080117",
	"TimezoneOffsetFromUTC":               "00080201",
	"NetworkID":                           "00081000",
	"StationName":                         "00081010",
	"StudyDescription":                    "00081030",
	"ProcedureCodeSequence":               "00081032",
	"SeriesDescription":                   "0008103E",
	"InstitutionalDepartmentName":         "00081040",
	"PhysiciansOfRecord":                  "00081048",
	"PhysiciansOfRecordIDSequence":        "00081049",
	"PerformingPhysicianName":             "00081050",
	"PerformingPhysicianIDSequence":       "00081052",
	"NameOfPhysicianReadingStudy":         "00081060",
	"PhysicianReadingStudyIDSequence":     "00081062",
	"OperatorsName":                       "00081070",
	"OperatorIDSequence":                  "00081072",
	"AdmittingDiagnosesDescription":       "00081080",
	"AdmittingDiagnosesCodeSequence":      "00081084",
	"ManufacturersModelName":              "00081090",
	"ReferencedResultsSequence":           "00081100",
	"ReferencedStudySequence":             "00081110",
	"ReferencedProcedureStepSequence":     "00081111",
	"ReferencedSeriesSequence":            "00081115",
	"ReferencedPatientSequence":           "00081120",
	"ReferencedVisitSequence":             "00081125",
	"ReferencedOverlaySequence":           "00081130",
	"ReferencedWaveformSequence":          "0008113A",
	"ReferencedImageSequence":             "00081140",
	"ReferencedCurveSequence":             "00081145",
	"ReferencedInstanceSequence":          "0008114A",
	"ReferencedSOPClassUID":               "00081150",
	"ReferencedSOPInstanceUID":            "00081155",
	"SOPClassesSupported":                 "0008115A",
	"ReferencedFrameNumber":               "00081160",
	"SimpleFrameList":                     "00081161",
	"CalculatedFrameList":                 "00081162",
	"TimeRange":                           "00081163",
	"FrameExtractionSequence":             "00081164",
	"RetrieveURL":                         "00081190",
	"TransactionUID":                      "00081195",
	"FailureReason":                       "00081197",
	"FailedSOPSequence":                   "00081198",
	"ReferencedSOPSequence":               "00081199",
	"OtherReferencedStudiesSequence":      "00081200",
	"RelatedSeriesSequence":               "00081250",
	"LossyImageCompression":               "00082110",
	"DerivationDescription":               "00082111",
	"SourceImageSequence":                 "00082112",
	"StageName":                           "00082120",
	"StageNumber":                         "00082122",
	"NumberOfStages":                      "00082124",
	"ViewName":                            "00082127",
	"ViewNumber":                          "00082128",
	"NumberOfEventTimers":                 "00082129",
	"NumberOfViewsInStage":                "0008212A",
	"EventElapsedTimes":                   "00082130",
	"EventTimerNames":                     "00082132",
	"EventTimerSequence":                  "00082133",
	"EventTimeOffset":                     "00082134",
	"EventCodeSequence":                   "00082135",
	"StartTrim":                           "00082142",
	"StopTrim":                            "00082143",
	"RecommendedDisplayFrameRate":         "00082144",
	"TransducerPosition":                  "00082200",
	"TransducerOrientation":               "00082204",
	"AnatomicStructure":                   "00082208",
	"AnatomicRegionSequence":              "00082218",
	"AnatomicRegionModifierSequence":      "00082220",
	"PrimaryAnatomicStructureSequence":    "00082228",
	"AnatomicStructureOrRegionSequence":   "00082229",
	"AnatomicStructureModifierSequence":   "00082230",
	"TransducerPositionSequence":          "00082240",
	"TransducerPositionModifierSequence":  "00082242",
	"TransducerOrientationSequence":       "00082244",
	"TransducerOrientationModifierSeq":    "00082246",
	"AnatomicEntrancePortalCodeSeqTrial":  "00082253",
	"AnatomicApproachDirCodeSeqTrial":     "00082255",
	"AnatomicPerspectiveDescrTrial":       "00082256",
	"AnatomicPerspectiveCodeSeqTrial":     "00082257",
	"AlternateRepresentationSequence":     "00083001",
	"IrradiationEventUID":                 "00083010",
	"IdentifyingComments":                 "00084000",
	"FrameType":                           "00089007",
	"ReferencedImageEvidenceSequence":     "00089092",
	"ReferencedRawDataSequence":           "00089121",
	"CreatorVersionUID":                   "00089123",
	"DerivationImageSequence":             "00089124",
	"SourceImageEvidenceSequence":         "00089154",
	"PixelPresentation":                   "00089205",
	"VolumetricProperties":                "00089206",
	"VolumeBasedCalculationTechnique":     "00089207",
	"ComplexImageComponent":               "00089208",
	"AcquisitionContrast":                 "00089209",
	"DerivationCodeSequence":              "00089215",
	"GrayscalePresentationStateSequence":  "00089237",
	"ReferencedOtherPlaneSequence":        "00089410",
	"FrameDisplaySequence":                "00089458",
	"RecommendedDisplayFrameRateInFloat":  "00089459",
	"SkipFrameRangeFlag":                  "00089460",
	"FullFidelity":                        "00091001",
	"SuiteID":                             "00091002",
	"ProductID":                           "00091004",
	"ImageActualDate":                     "00091027",
	"ServiceID":                           "00091030",
	"MobileLocationNumber":                "00091031",
	"EquipmentUID":                        "000910E3",
	"GenesisVersionNow":                   "000910E6",
	"ExamRecordChecksum":                  "000910E7",
	"ActualSeriesDataTimeStamp":           "000910E9",
	"PatientGroupLength":                  "00100000",
	"PatientName":                         "00100010",
	"PatientID":                           "00100020",
	"IssuerOfPatientID":                   "00100021",
	"TypeOfPatientID":                     "00100022",
	"PatientBirthDate":                    "00100030",
	"PatientBirthTime":                    "00100032",
	"PatientSex":                          "00100040",
	"PatientInsurancePlanCodeSequence":    "00100050",
	"PatientPrimaryLanguageCodeSeq":       "00100101",
	"PatientPrimaryLanguageCodeModSeq":    "00100102",
	"OtherPatientIDs":                     "00101000",
	"OtherPatientNames":                   "00101001",
	"OtherPatientIDsSequence":             "00101002",
	"PatientBirthName":                    "00101005",
	"PatientAge":                          "00101010",
	"PatientSize":                         "00101020",
	"PatientWeight":                       "00101030",
	"PatientAddress":                      "00101040",
	"InsurancePlanIdentification":         "00101050",
	"PatientMotherBirthName":              "00101060",
	"MilitaryRank":                        "00101080",
	"BranchOfService":                     "00101081",
	"MedicalRecordLocator":                "00101090",
	"MedicalAlerts":                       "00102000",
	"Allergies":                           "00102110",
	"CountryOfResidence":                  "00102150",
	"RegionOfResidence":                   "00102152",
	"PatientTelephoneNumbers":             "00102154",
	"EthnicGroup":                         "00102160",
	"Occupation":                          "00102180",
	"SmokingStatus":                       "001021A0",
	"AdditionalPatientHistory":            "001021B0",
	"PregnancyStatus":                     "001021C0",
	"LastMenstrualDate":                   "001021D0",
	"PatientReligiousPreference":          "001021F0",
	"PatientSpeciesDescription":           "00102201",
	"PatientSpeciesCodeSequence":          "00102202",
	"PatientSexNeutered":                  "00102203",
	"AnatomicalOrientationType":           "00102210",
	"PatientBreedDescription":             "00102292",
	"PatientBreedCodeSequence":            "00102293",
	"BreedRegistrationSequence":           "00102294",
	"BreedRegistrationNumber":             "00102295",
	"BreedRegistryCodeSequence":           "00102296",
	"ResponsiblePerson":                   "00102297",
	"ResponsiblePersonRole":               "00102298",
	"ResponsibleOrganization":             "00102299",
	"PatientComments":                     "00104000",
	"ExaminedBodyThickness":               "00109431",
	"PatientStatus":                       "00111010",
	"ClinicalTrialSponsorName":            "00120010",
	"ClinicalTrialProtocolID":             "00120020",
	"ClinicalTrialProtocolName":           "00120021",
	"ClinicalTrialSiteID":                 "00120030",
	"ClinicalTrialSiteName":               "00120031",
	"ClinicalTrialSubjectID":              "00120040",
	"ClinicalTrialSubjectReadingID":       "00120042",
	"ClinicalTrialTimePointID":            "00120050",
	"ClinicalTrialTimePointDescription":   "00120051",
	"ClinicalTrialCoordinatingCenter":     "00120060",
	"PatientIdentityRemoved":              "00120062",
	"DeidentificationMethod":              "00120063",
	"DeidentificationMethodCodeSequence":  "00120064",
	"ClinicalTrialSeriesID":               "00120071",
	"ClinicalTrialSeriesDescription":      "00120072",
	"DistributionType":                    "00120084",
	"ConsentForDistributionFlag":          "00120085",
	"AcquisitionGroupLength":              "00180000",
	"ContrastBolusAgent":                  "00180010",
	"ContrastBolusAgentSequence":          "00180012",
	"ContrastBolusAdministrationRoute":    "00180014",
	"BodyPartExamined":                    "00180015",
	"ScanningSequence":                    "00180020",
	"SequenceVariant":                     "00180021",
	"ScanOptions":                         "00180022",
	"MRAcquisitionType":                   "00180023",
	"SequenceName":                        "00180024",
	"AngioFlag":                           "00180025",
	"InterventionDrugInformationSeq":      "00180026",
	"InterventionDrugStopTime":            "00180027",
	"InterventionDrugDose":                "00180028",
	"InterventionDrugSequence":            "00180029",
	"AdditionalDrugSequence":              "0018002A",
	"Radionuclide":                        "00180030",
	"Radiopharmaceutical":                 "00180031",
	"EnergyWindowCenterline":              "00180032",
	"EnergyWindowTotalWidth":              "00180033",
	"InterventionDrugName":                "00180034",
	"InterventionDrugStartTime":           "00180035",
	"InterventionSequence":                "00180036",
	"TherapyType":                         "00180037",
	"InterventionStatus":                  "00180038",
	"TherapyDescription":                  "00180039",
	"InterventionDescription":             "0018003A",
	"CineRate":                            "00180040",
	"InitialCineRunState":                 "00180042",
	"SliceThickness":                      "00180050",
	"KVP":                                 "00180060",
	"CountsAccumulated":                   "00180070",
	"AcquisitionTerminationCondition":     "00180071",
	"EffectiveDuration":                   "00180072",
	"AcquisitionStartCondition":           "00180073",
	"AcquisitionStartConditionData":       "00180074",
	"AcquisitionEndConditionData":         "00180075",
	"RepetitionTime":                      "00180080",
	"EchoTime":                            "00180081",
	"InversionTime":                       "00180082",
	"NumberOfAverages":                    "00180083",
	"ImagingFrequency":                    "00180084",
	"ImagedNucleus":                       "00180085",
	"EchoNumber":                          "00180086",
	"MagneticFieldStrength":               "00180087",
	"SpacingBetweenSlices":                "00180088",
	"NumberOfPhaseEncodingSteps":          "00180089",
	"DataCollectionDiameter":              "00180090",
	"EchoTrainLength":                     "00180091",
	"PercentSampling":                     "00180093",
	"PercentPhaseFieldOfView":             "00180094",
	"PixelBandwidth":                      "00180095",
	"DeviceSerialNumber":                  "00181000",
	"DeviceUID":                           "00181002",
	"DeviceID":                            "00181003",
	"PlateID":                             "00181004",
	"GeneratorID":                         "00181005",
	"GridID":                              "00181006",
	"CassetteID":                          "00181007",
	"GantryID":                            "00181008",
	"SecondaryCaptureDeviceID":            "00181010",
	"HardcopyCreationDeviceID":            "00181011",
	"DateOfSecondaryCapture":              "00181012",
	"TimeOfSecondaryCapture":              "00181014",
	"SecondaryCaptureDeviceManufacturer":  "00181016",
	"HardcopyDeviceManufacturer":          "00181017",
	"SecondaryCaptureDeviceModelName":     "00181018",
	"SecondaryCaptureDeviceSoftwareVers":  "00181019",
	"HardcopyDeviceSoftwareVersion":       "0018101A",
	"HardcopyDeviceModelName":             "0018101B",
	"SoftwareVersion":                     "00181020",
	"VideoImageFormatAcquired":            "00181022",
	"DigitalImageFormatAcquired":          "00181023",
	"ProtocolName":                        "00181030",
	"ContrastBolusRoute":                  "00181040",
	"ContrastBolusVolume":                 "00181041",
	"ContrastBolusStartTime":              "00181042",
	"ContrastBolusStopTime":               "00181043",
	"ContrastBolusTotalDose":              "00181044",
	"SyringeCounts":                       "00181045",
	"ContrastFlowRate":                    "00181046",
	"ContrastFlowDuration":                "00181047",
	"ContrastBolusIngredient":             "00181048",
	"ContrastBolusConcentration":          "00181049",
	"SpatialResolution":                   "00181050",
	"TriggerTime":                         "00181060",
	"TriggerSourceOrType":                 "00181061",
	"NominalInterval":                     "00181062",
	"FrameTime":                           "00181063",
	"CardiacFramingType":                  "00181064",
	"FrameTimeVector":                     "00181065",
	"FrameDelay":                          "00181066",
	"ImageTriggerDelay":                   "00181067",
	"MultiplexGroupTimeOffset":            "00181068",
	"TriggerTimeOffset":                   "00181069",
	"SynchronizationTrigger":              "0018106A",
	"SynchronizationChannel":              "0018106C",
	"TriggerSamplePosition":               "0018106E",
	"RadiopharmaceuticalRoute":            "00181070",
	"RadiopharmaceuticalVolume":           "00181071",
	"RadiopharmaceuticalStartTime":        "00181072",
	"RadiopharmaceuticalStopTime":         "00181073",
	"RadionuclideTotalDose":               "00181074",
	"RadionuclideHalfLife":                "00181075",
	"RadionuclidePositronFraction":        "00181076",
	"RadiopharmaceuticalSpecActivity":     "00181077",
	"RadiopharmaceuticalStartDateTime":    "00181078",
	"RadiopharmaceuticalStopDateTime":     "00181079",
	"BeatRejectionFlag":                   "00181080",
	"LowRRValue":                          "00181081",
	"HighRRValue":                         "00181082",
	"IntervalsAcquired":                   "00181083",
	"IntervalsRejected":                   "00181084",
	"PVCRejection":                        "00181085",
	"SkipBeats":                           "00181086",
	"HeartRate":                           "00181088",
	"CardiacNumberOfImages":               "00181090",
	"TriggerWindow":                       "00181094",
	"ReconstructionDiameter":              "00181100",
	"DistanceSourceToDetector":            "00181110",
	"DistanceSourceToPatient":             "00181111",
	"EstimatedRadiographicMagnification":  "00181114",
	"GantryDetectorTilt":                  "00181120",
	"GantryDetectorSlew":                  "00181121",
	"TableHeight":                         "00181130",
	"TableTraverse":                       "00181131",
	"TableMotion":                         "00181134",
	"TableVerticalIncrement":              "00181135",
	"TableLateralIncrement":               "00181136",
	"TableLongitudinalIncrement":          "00181137",
	"TableAngle":                          "00181138",
	"TableType":                           "0018113A",
	"RotationDirection":                   "00181140",
	"AngularPosition":                     "00181141",
	"RadialPosition":                      "00181142",
	"ScanArc":                             "00181143",
	"AngularStep":                         "00181144",
	"CenterOfRotationOffset":              "00181145",
	"RotationOffset":                      "00181146",
	"FieldOfViewShape":                    "00181147",
	"FieldOfViewDimensions":               "00181149",
	"ExposureTime":                        "00181150",
	"XRayTubeCurrent":                     "00181151",
	"Exposure":                            "00181152",
	"ExposureInMicroAmpSec":               "00181153",
	"AveragePulseWidth":                   "00181154",
	"RadiationSetting":                    "00181155",
	"RectificationType":                   "00181156",
	"RadiationMode":                       "0018115A",
	"ImageAreaDoseProduct":                "0018115E",
	"FilterType":                          "00181160",
	"TypeOfFilters":                       "00181161",
	"IntensifierSize":                     "00181162",
	"ImagerPixelSpacing":                  "00181164",
	"Grid":                                "00181166",
	"GeneratorPower":                      "00181170",
	"CollimatorGridName":                  "00181180",
	"CollimatorType":                      "00181181",
	"FocalDistance":                       "00181182",
	"XFocusCenter":                        "00181183",
	"YFocusCenter":                        "00181184",
	"FocalSpots":                          "00181190",
	"AnodeTargetMaterial":                 "00181191",
	"BodyPartThickness":                   "001811A0",
	"CompressionForce":                    "001811A2",
	"DateOfLastCalibration":               "00181200",
	"TimeOfLastCalibration":               "00181201",
	"ConvolutionKernel":                   "00181210",
	"UpperLowerPixelValues":               "00181240",
	"ActualFrameDuration":                 "00181242",
	"CountRate":                           "00181243",
	"PreferredPlaybackSequencing":         "00181244",
	"ReceiveCoilName":                     "00181250",
	"TransmitCoilName":                    "00181251",
	"PlateType":                           "00181260",
	"PhosphorType":                        "00181261",
	"ScanVelocity":                        "00181300",
	"WholeBodyTechnique":                  "00181301",
	"ScanLength":                          "00181302",
	"AcquisitionMatrix":                   "00181310",
	"InPlanePhaseEncodingDirection":       "00181312",
	"FlipAngle":                           "00181314",
	"VariableFlipAngleFlag":               "00181315",
	"SAR":                                 "00181316",
	"DBDt":                                "00181318",
	"AcquisitionDeviceProcessingDescr":    "00181400",
	"AcquisitionDeviceProcessingCode":     "00181401",
	"CassetteOrientation":                 "00181402",
	"CassetteSize":                        "00181403",
	"ExposuresOnPlate":                    "00181404",
	"RelativeXRayExposure":                "00181405",
	"ColumnAngulation":                    "00181450",
	"TomoLayerHeight":                     "00181460",
	"TomoAngle":                           "00181470",
	"TomoTime":                            "00181480",
	"TomoType":                            "00181490",
	"TomoClass":                           "00181491",
	"NumberOfTomosynthesisSourceImages":   "00181495",
	"PositionerMotion":                    "00181500",
	"PositionerType":                      "00181508",
	"PositionerPrimaryAngle":              "00181510",
	"PositionerSecondaryAngle":            "00181511",
	"PositionerPrimaryAngleIncrement":     "00181520",
	"PositionerSecondaryAngleIncrement":   "00181521",
	"DetectorPrimaryAngle":                "00181530",
	"DetectorSecondaryAngle":              "00181531",
	"ShutterShape":                        "00181600",
	"ShutterLeftVerticalEdge":             "00181602",
	"ShutterRightVerticalEdge":            "00181604",
	"ShutterUpperHorizontalEdge":          "00181606",
	"ShutterLowerHorizontalEdge":          "00181608",
	"CenterOfCircularShutter":             "00181610",
	"RadiusOfCircularShutter":             "00181612",
	"VerticesOfPolygonalShutter":          "00181620",
	"ShutterPresentationValue":            "00181622",
	"ShutterOverlayGroup":                 "00181623",
	"ShutterPresentationColorCIELabVal":   "00181624",
	"CollimatorShape":                     "00181700",
	"CollimatorLeftVerticalEdge":          "00181702",
	"CollimatorRightVerticalEdge":         "00181704",
	"CollimatorUpperHorizontalEdge":       "00181706",
	"CollimatorLowerHorizontalEdge":       "00181708",
	"CenterOfCircularCollimator":          "00181710",
	"RadiusOfCircularCollimator":          "00181712",
	"VerticesOfPolygonalCollimator":       "00181720",
	"AcquisitionTimeSynchronized":         "00181800",
	"TimeSource":                          "00181801",
	"TimeDistributionProtocol":            "00181802",
	"NTPSourceAddress":                    "00181803",
	"PageNumberVector":                    "00182001",
	"FrameLabelVector":                    "00182002",
	"FramePrimaryAngleVector":             "00182003",
	"FrameSecondaryAngleVector":           "00182004",
	"SliceLocationVector":                 "00182005",
	"DisplayWindowLabelVector":            "00182006",
	"NominalScannedPixelSpacing":          "00182010",
	"DigitizingDeviceTransportDirection":  "00182020",
	"RotationOfScannedFilm":               "00182030",
	"IVUSAcquisition":                     "00183100",
	"IVUSPullbackRate":                    "00183101",
	"IVUSGatedRate":                       "00183102",
	"IVUSPullbackStartFrameNumber":        "00183103",
	"IVUSPullbackStopFrameNumber":         "00183104",
	"LesionNumber":                        "00183105",
	"AcquisitionComments":                 "00184000",
	"OutputPower":                         "00185000",
	"TransducerData":                      "00185010",
	"FocusDepth":                          "00185012",
	"ProcessingFunction":                  "00185020",
	"PostprocessingFunction":              "00185021",
	"MechanicalIndex":                     "00185022",
	"BoneThermalIndex":                    "00185024",
	"CranialThermalIndex":                 "00185026",
	"SoftTissueThermalIndex":              "00185027",
	"SoftTissueFocusThermalIndex":         "00185028",
	"SoftTissueSurfaceThermalIndex":       "00185029",
	"DynamicRange":                        "00185030",
	"TotalGain":                           "00185040",
	"DepthOfScanField":                    "00185050",
	"PatientPosition":                     "00185100",
	"ViewPosition":                        "00185101",
	"ProjectionEponymousNameCodeSeq":      "00185104",
	"ImageTransformationMatrix":           "00185210",
	"ImageTranslationVector":              "00185212",
	"Sensitivity":                         "00186000",
	"SequenceOfUltrasoundRegions":         "00186011",
	"RegionSpatialFormat":                 "00186012",
	"RegionDataType":                      "00186014",
	"RegionFlags":                         "00186016",
	"RegionLocationMinX0":                 "00186018",
	"RegionLocationMinY0":                 "0018601A",
	"RegionLocationMaxX1":                 "0018601C",
	"RegionLocationMaxY1":                 "0018601E",
	"ReferencePixelX0":                    "00186020",
	"ReferencePixelY0":                    "00186022",
	"PhysicalUnitsXDirection":             "00186024",
	"PhysicalUnitsYDirection":             "00186026",
	"ReferencePixelPhysicalValueX":        "00186028",
	"ReferencePixelPhysicalValueY":        "0018602A",
	"PhysicalDeltaX":                      "0018602C",
	"PhysicalDeltaY":                      "0018602E",
	"TransducerFrequency":                 "00186030",
	"TransducerType":                      "00186031",
	"PulseRepetitionFrequency":            "00186032",
	"DopplerCorrectionAngle":              "00186034",
	"SteeringAngle":                       "00186036",
	"DopplerSampleVolumeXPosRetired":      "00186038",
	"DopplerSampleVolumeXPosition":        "00186039",
	"DopplerSampleVolumeYPosRetired":      "0018603A",
	"DopplerSampleVolumeYPosition":        "0018603B",
	"TMLinePositionX0Retired":             "0018603C",
	"TMLinePositionX0":                    "0018603D",
	"TMLinePositionY0Retired":             "0018603E",
	"TMLinePositionY0":                    "0018603F",
	"TMLinePositionX1Retired":             "00186040",
	"TMLinePositionX1":                    "00186041",
	"TMLinePositionY1Retired":             "00186042",
	"TMLinePositionY1":                    "00186043",
	"PixelComponentOrganization":          "00186044",
	"PixelComponentMask":                  "00186046",
	"PixelComponentRangeStart":            "00186048",
	"PixelComponentRangeStop":             "0018604A",
	"PixelComponentPhysicalUnits":         "0018604C",
	"PixelComponentDataType":              "0018604E",
	"NumberOfTableBreakPoints":            "00186050",
	"TableOfXBreakPoints":                 "00186052",
	"TableOfYBreakPoints":                 "00186054",
	"NumberOfTableEntries":                "00186056",
	"TableOfPixelValues":                  "00186058",
	"TableOfParameterValues":              "0018605A",
	"RWaveTimeVector":                     "00186060",
	"DetectorConditionsNominalFlag":       "00187000",
	"DetectorTemperature":                 "00187001",
	"DetectorType":                        "00187004",
	"DetectorConfiguration":               "00187005",
	"DetectorDescription":                 "00187006",
	"DetectorMode":                        "00187008",
	"DetectorID":                          "0018700A",
	"DateOfLastDetectorCalibration":       "0018700C",
	"TimeOfLastDetectorCalibration":       "0018700E",
	"DetectorExposuresSinceCalibration":   "00187010",
	"DetectorExposuresSinceManufactured":  "00187011",
	"DetectorTimeSinceLastExposure":       "00187012",
	"DetectorActiveTime":                  "00187014",
	"DetectorActiveOffsetFromExposure":    "00187016",
	"DetectorBinning":                     "0018701A",
	"DetectorElementPhysicalSize":         "00187020",
	"DetectorElementSpacing":              "00187022",
	"DetectorActiveShape":                 "00187024",
	"DetectorActiveDimensions":            "00187026",
	"DetectorActiveOrigin":                "00187028",
	"DetectorManufacturerName":            "0018702A",
	"DetectorManufacturersModelName":      "0018702B",
	"FieldOfViewOrigin":                   "00187030",
	"FieldOfViewRotation":                 "00187032",
	"FieldOfViewHorizontalFlip":           "00187034",
	"GridAbsorbingMaterial":               "00187040",
	"GridSpacingMaterial":                 "00187041",
	"GridThickness":                       "00187042",
	"GridPitch":                           "00187044",
	"GridAspectRatio":                     "00187046",
	"GridPeriod":                          "00187048",
	"GridFocalDistance":                   "0018704C",
	"FilterMaterial":                      "00187050",
	"FilterThicknessMinimum":              "00187052",
	"FilterThicknessMaximum":              "00187054",
	"ExposureControlMode":                 "00187060",
	"ExposureControlModeDescription":      "00187062",
	"ExposureStatus":                      "00187064",
	"PhototimerSetting":                   "00187065",
	"ExposureTimeInMicroSec":              "00188150",
	"XRayTubeCurrentInMicroAmps":          "00188151",
	"ContentQualification":                "00189004",
	"PulseSequenceName":                   "00189005",
	"MRImagingModifierSequence":           "00189006",
	"EchoPulseSequence":                   "00189008",
	"InversionRecovery":                   "00189009",
	"FlowCompensation":                    "00189010",
	"MultipleSpinEcho":                    "00189011",
	"MultiPlanarExcitation":               "00189012",
	"PhaseContrast":                       "00189014",
	"TimeOfFlightContrast":                "00189015",
	"Spoiling":                            "00189016",
	"SteadyStatePulseSequence":            "00189017",
	"EchoPlanarPulseSequence":             "00189018",
	"TagAngleFirstAxis":                   "00189019",
	"MagnetizationTransfer":               "00189020",
	"T2Preparation":                       "00189021",
	"BloodSignalNulling":                  "00189022",
	"SaturationRecovery":                  "00189024",
	"SpectrallySelectedSuppression":       "00189025",
	"SpectrallySelectedExcitation":        "00189026",
	"SpatialPresaturation":                "00189027",
	"Tagging":                             "00189028",
	"OversamplingPhase":                   "00189029",
	"TagSpacingFirstDimension":            "00189030",
	"GeometryOfKSpaceTraversal":           "00189032",
	"SegmentedKSpaceTraversal":            "00189033",
	"RectilinearPhaseEncodeReordering":    "00189034",
	"TagThickness":                        "00189035",
	"PartialFourierDirection":             "00189036",
	"CardiacSynchronizationTechnique":     "00189037",
	"ReceiveCoilManufacturerName":         "00189041",
	"MRReceiveCoilSequence":               "00189042",
	"ReceiveCoilType":                     "00189043",
	"QuadratureReceiveCoil":               "00189044",
	"MultiCoilDefinitionSequence":         "00189045",
	"MultiCoilConfiguration":              "00189046",
	"MultiCoilElementName":                "00189047",
	"MultiCoilElementUsed":                "00189048",
	"MRTransmitCoilSequence":              "00189049",
	"TransmitCoilManufacturerName":        "00189050",
	"TransmitCoilType":                    "00189051",
	"SpectralWidth":                       "00189052",
	"ChemicalShiftReference":              "00189053",
	"VolumeLocalizationTechnique":         "00189054",
	"MRAcquisitionFrequencyEncodeSteps":   "00189058",
	"Decoupling":                          "00189059",
	"DecoupledNucleus":                    "00189060",
	"DecouplingFrequency":                 "00189061",
	"DecouplingMethod":                    "00189062",
	"DecouplingChemicalShiftReference":    "00189063",
	"KSpaceFiltering":                     "00189064",
	"TimeDomainFiltering":                 "00189065",
	"NumberOfZeroFills":                   "00189066",
	"BaselineCorrection":                  "00189067",
	"ParallelReductionFactorInPlane":      "00189069",
	"CardiacRRIntervalSpecified":          "00189070",
	"AcquisitionDuration":                 "00189073",
	"FrameAcquisitionDateTime":            "00189074",
	"DiffusionDirectionality":             "00189075",
	"DiffusionGradientDirectionSequence":  "00189076",
	"ParallelAcquisition":                 "00189077",
	"ParallelAcquisitionTechnique":        "00189078",
	"InversionTimes":                      "00189079",
	"MetaboliteMapDescription":            "00189080",
	"PartialFourier":                      "00189081",
	"EffectiveEchoTime":                   "00189082",
	"MetaboliteMapCodeSequence":           "00189083",
	"ChemicalShiftSequence":               "00189084",
	"CardiacSignalSource":                 "00189085",
	"DiffusionBValue":                     "00189087",
	"DiffusionGradientOrientation":        "00189089",
	"VelocityEncodingDirection":           "00189090",
	"VelocityEncodingMinimumValue":        "00189091",
	"NumberOfKSpaceTrajectories":          "00189093",
	"CoverageOfKSpace":                    "00189094",
	"SpectroscopyAcquisitionPhaseRows":    "00189095",
	"ParallelReductFactorInPlaneRetired":  "00189096",
	"TransmitterFrequency":                "00189098",
	"ResonantNucleus":                     "00189100",
	"FrequencyCorrection":                 "00189101",
	"MRSpectroscopyFOVGeometrySequence":   "00189103",
	"SlabThickness":                       "00189104",
	"SlabOrientation":                     "00189105",
	"MidSlabPosition":                     "00189106",
	"MRSpatialSaturationSequence":         "00189107",
	"MRTimingAndRelatedParametersSeq":     "00189112",
	"MREchoSequence":                      "00189114",
	"MRModifierSequence":                  "00189115",
	"MRDiffusionSequence":                 "00189117",
	"CardiacTriggerSequence":              "00189118",
	"MRAveragesSequence":                  "00189119",
	"MRFOVGeometrySequence":               "00189125",
	"VolumeLocalizationSequence":          "00189126",
	"SpectroscopyAcquisitionDataColumns":  "00189127",
	"DiffusionAnisotropyType":             "00189147",
	"FrameReferenceDateTime":              "00189151",
	"MRMetaboliteMapSequence":             "00189152",
	"ParallelReductionFactorOutOfPlane":   "00189155",
	"SpectroscopyOutOfPlanePhaseSteps":    "00189159",
	"BulkMotionStatus":                    "00189166",
	"ParallelReductionFactSecondInPlane":  "00189168",
	"CardiacBeatRejectionTechnique":       "00189169",
	"RespiratoryMotionCompTechnique":      "00189170",
	"RespiratorySignalSource":             "00189171",
	"BulkMotionCompensationTechnique":     "00189172",
	"BulkMotionSignalSource":              "00189173",
	"ApplicableSafetyStandardAgency":      "00189174",
	"ApplicableSafetyStandardDescr":       "00189175",
	"OperatingModeSequence":               "00189176",
	"OperatingModeType":                   "00189177",
	"OperatingMode":                       "00189178",
	"SpecificAbsorptionRateDefinition":    "00189179",
	"GradientOutputType":                  "00189180",
	"SpecificAbsorptionRateValue":         "00189181",
	"GradientOutput":                      "00189182",
	"FlowCompensationDirection":           "00189183",
	"TaggingDelay":                        "00189184",
	"RespiratoryMotionCompTechDescr":      "00189185",
	"RespiratorySignalSourceID":           "00189186",
	"ChemicalShiftsMinIntegrateLimitHz":   "00189195",
	"ChemicalShiftsMaxIntegrateLimitHz":   "00189196",
	"MRVelocityEncodingSequence":          "00189197",
	"FirstOrderPhaseCorrection":           "00189198",
	"WaterReferencedPhaseCorrection":      "00189199",
	"MRSpectroscopyAcquisitionType":       "00189200",
	"RespiratoryCyclePosition":            "00189214",
	"VelocityEncodingMaximumValue":        "00189217",
	"TagSpacingSecondDimension":           "00189218",
	"TagAngleSecondAxis":                  "00189219",
	"FrameAcquisitionDuration":            "00189220",
	"MRImageFrameTypeSequence":            "00189226",
	"MRSpectroscopyFrameTypeSequence":     "00189227",
	"MRAcqPhaseEncodingStepsInPlane":      "00189231",
	"MRAcqPhaseEncodingStepsOutOfPlane":   "00189232",
	"SpectroscopyAcqPhaseColumns":         "00189234",
	"CardiacCyclePosition":                "00189236",
	"SpecificAbsorptionRateSequence":      "00189239",
	"RFEchoTrainLength":                   "00189240",
	"GradientEchoTrainLength":             "00189241",
	"ChemicalShiftsMinIntegrateLimitPPM":  "00189295",
	"ChemicalShiftsMaxIntegrateLimitPPM":  "00189296",
	"CTAcquisitionTypeSequence":           "00189301",
	"AcquisitionType":                     "00189302",
	"TubeAngle":                           "00189303",
	"CTAcquisitionDetailsSequence":        "00189304",
	"RevolutionTime":                      "00189305",
	"SingleCollimationWidth":              "00189306",
	"TotalCollimationWidth":               "00189307",
	"CTTableDynamicsSequence":             "00189308",
	"TableSpeed":                          "00189309",
	"TableFeedPerRotation":                "00189310",
	"SpiralPitchFactor":                   "00189311",
	"CTGeometrySequence":                  "00189312",
	"DataCollectionCenterPatient":         "00189313",
	"CTReconstructionSequence":            "00189314",
	"ReconstructionAlgorithm":             "00189315",
	"ConvolutionKernelGroup":              "00189316",
	"ReconstructionFieldOfView":           "00189317",
	"ReconstructionTargetCenterPatient":   "00189318",
	"ReconstructionAngle":                 "00189319",
	"ImageFilter":                         "00189320",
	"CTExposureSequence":                  "00189321",
	"ReconstructionPixelSpacing":          "00189322",
	"ExposureModulationType":              "00189323",
	"EstimatedDoseSaving":                 "00189324",
	"CTXRayDetailsSequence":               "00189325",
	"CTPositionSequence":                  "00189326",
	"TablePosition":                       "00189327",
	"ExposureTimeInMilliSec":              "00189328",
	"CTImageFrameTypeSequence":            "00189329",
	"XRayTubeCurrentInMilliAmps":          "00189330",
	"ExposureInMilliAmpSec":               "00189332",
	"ConstantVolumeFlag":                  "00189333",
	"FluoroscopyFlag":                     "00189334",
	"SourceToDataCollectionCenterDist":    "00189335",
	"ContrastBolusAgentNumber":            "00189337",
	"ContrastBolusIngredientCodeSeq":      "00189338",
	"ContrastAdministrationProfileSeq":    "00189340",
	"ContrastBolusUsageSequence":          "00189341",
	"ContrastBolusAgentAdministered":      "00189342",
	"ContrastBolusAgentDetected":          "00189343",
	"ContrastBolusAgentPhase":             "00189344",
	"CTDIvol":                             "00189345",
	"CTDIPhantomTypeCodeSequence":         "00189346",
	"CalciumScoringMassFactorPatient":     "00189351",
	"CalciumScoringMassFactorDevice":      "00189352",
	"EnergyWeightingFactor":               "00189353",
	"CTAdditionalXRaySourceSequence":      "00189360",
	"ProjectionPixelCalibrationSequence":  "00189401",
	"DistanceSourceToIsocenter":           "00189402",
	"DistanceObjectToTableTop":            "00189403",
	"ObjectPixelSpacingInCenterOfBeam":    "00189404",
	"PositionerPositionSequence":          "00189405",
	"TablePositionSequence":               "00189406",
	"CollimatorShapeSequence":             "00189407",
	"XAXRFFrameCharacteristicsSequence":   "00189412",
	"FrameAcquisitionSequence":            "00189417",
	"XRayReceptorType":                    "00189420",
	"AcquisitionProtocolName":             "00189423",
	"AcquisitionProtocolDescription":      "00189424",
	"ContrastBolusIngredientOpaque":       "00189425",
	"DistanceReceptorPlaneToDetHousing":   "00189426",
	"IntensifierActiveShape":              "00189427",
	"IntensifierActiveDimensions":         "00189428",
	"PhysicalDetectorSize":                "00189429",
	"PositionOfIsocenterProjection":       "00189430",
	"FieldOfViewSequence":                 "00189432",
	"FieldOfViewDescription":              "00189433",
	"ExposureControlSensingRegionsSeq":    "00189434",
	"ExposureControlSensingRegionShape":   "00189435",
	"ExposureControlSensRegionLeftEdge":   "00189436",
	"ExposureControlSensRegionRightEdge":  "00189437",
	"CenterOfCircExposControlSensRegion":  "00189440",
	"RadiusOfCircExposControlSensRegion":  "00189441",
	"ColumnAngulationPatient":             "00189447",
	"BeamAngle":                           "00189449",
	"FrameDetectorParametersSequence":     "00189451",
	"CalculatedAnatomyThickness":          "00189452",
	"CalibrationSequence":                 "00189455",
	"ObjectThicknessSequence":             "00189456",
	"PlaneIdentification":                 "00189457",
	"FieldOfViewDimensionsInFloat":        "00189461",
	"IsocenterReferenceSystemSequence":    "00189462",
	"PositionerIsocenterPrimaryAngle":     "00189463",
	"PositionerIsocenterSecondaryAngle":   "00189464",
	"PositionerIsocenterDetRotAngle":      "00189465",
	"TableXPositionToIsocenter":           "00189466",
	"TableYPositionToIsocenter":           "00189467",
	"TableZPositionToIsocenter":           "00189468",
	"TableHorizontalRotationAngle":        "00189469",
	"TableHeadTiltAngle":                  "00189470",
	"TableCradleTiltAngle":                "00189471",
	"FrameDisplayShutterSequence":         "00189472",
	"AcquiredImageAreaDoseProduct":        "00189473",
	"CArmPositionerTabletopRelationship":  "00189474",
	"XRayGeometrySequence":                "00189476",
	"IrradiationEventIDSequence":          "00189477",
	"XRay3DFrameTypeSequence":             "00189504",
	"ContributingSourcesSequence":         "00189506",
	"XRay3DAcquisitionSequence":           "00189507",
	"PrimaryPositionerScanArc":            "00189508",
	"SecondaryPositionerScanArc":          "00189509",
	"PrimaryPositionerScanStartAngle":     "00189510",
	"SecondaryPositionerScanStartAngle":   "00189511",
	"PrimaryPositionerIncrement":          "00189514",
	"SecondaryPositionerIncrement":        "00189515",
	"StartAcquisitionDateTime":            "00189516",
	"EndAcquisitionDateTime":              "00189517",
	"ApplicationName":                     "00189524",
	"ApplicationVersion":                  "00189525",
	"ApplicationManufacturer":             "00189526",
	"AlgorithmType":                       "00189527",
	"AlgorithmDescription":                "00189528",
	"XRay3DReconstructionSequence":        "00189530",
	"ReconstructionDescription":           "00189531",
	"PerProjectionAcquisitionSequence":    "00189538",
	"DiffusionBMatrixSequence":            "00189601",
	"DiffusionBValueXX":                   "00189602",
	"DiffusionBValueXY":                   "00189603",
	"DiffusionBValueXZ":                   "00189604",
	"DiffusionBValueYY":                   "00189605",
	"DiffusionBValueYZ":                   "00189606",
	"DiffusionBValueZZ":                   "00189607",
	"DecayCorrectionDateTime":             "00189701",
	"StartDensityThreshold":               "00189715",
	"TerminationTimeThreshold":            "00189722",
	"DetectorGeometry":                    "00189725",
	"AxialDetectorDimension":              "00189727",
	"PETPositionSequence":                 "00189735",
	"NumberOfIterations":                  "00189739",
	"NumberOfSubsets":                     "00189740",
	"PETFrameTypeSequence":                "00189751",
	"ReconstructionType":                  "00189756",
	"DecayCorrected":                      "00189758",
	"AttenuationCorrected":                "00189759",
	"ScatterCorrected":                    "00189760",
	"DeadTimeCorrected":                   "00189761",
	"GantryMotionCorrected":               "00189762",
	"PatientMotionCorrected":              "00189763",
	"RandomsCorrected":                    "00189765",
	"SensitivityCalibrated":               "00189767",
	"DepthsOfFocus":                       "00189801",
	"ExclusionStartDatetime":              "00189804",
	"ExclusionDuration":                   "00189805",
	"ImageDataTypeSequence":               "00189807",
	"DataType":                            "00189808",
	"AliasedDataType":                     "0018980B",
	"ContributingEquipmentSequence":       "0018A001",
	"ContributionDateTime":                "0018A002",
	"ContributionDescription":             "0018A003",
	"NumberOfCellsIInDetector":            "00191002",
	"CellNumberAtTheta":                   "00191003",
	"CellSpacing":                         "00191004",
	"HorizFrameOfRef":                     "0019100F",
	"SeriesContrast":                      "00191011",
	"LastPseq":                            "00191012",
	"StartNumberForBaseline":              "00191013",
	"EndNumberForBaseline":                "00191014",
	"StartNumberForEnhancedScans":         "00191015",
	"EndNumberForEnhancedScans":           "00191016",
	"SeriesPlane":                         "00191017",
	"FirstScanRas":                        "00191018",
	"FirstScanLocation":                   "00191019",
	"LastScanRas":                         "0019101A",
	"LastScanLoc":                         "0019101B",
	"DisplayFieldOfView":                  "0019101E",
	"MidScanTime":                         "00191024",
	"MidScanFlag":                         "00191025",
	"DegreesOfAzimuth":                    "00191026",
	"GantryPeriod":                        "00191027",
	"XRayOnPosition":                      "0019102A",
	"XRayOffPosition":                     "0019102B",
	"NumberOfTriggers":                    "0019102C",
	"AngleOfFirstView":                    "0019102E",
	"TriggerFrequency":                    "0019102F",
	"ScanFOVType":                         "00191039",
	"StatReconFlag":                       "00191040",
	"ComputeType":                         "00191041",
	"SegmentNumber":                       "00191042",
	"TotalSegmentsRequested":              "00191043",
	"InterscanDelay":                      "00191044",
	"ViewCompressionFactor":               "00191047",
	"TotalNoOfRefChannels":                "0019104A",
	"DataSizeForScanData":                 "0019104B",
	"ReconPostProcflag":                   "00191052",
	"CTWaterNumber":                       "00191057",
	"CTBoneNumber":                        "00191058",
	"NumberOfChannels":                    "0019105E",
	"IncrementBetweenChannels":            "0019105F",
	"StartingView":                        "00191060",
	"NumberOfViews":                       "00191061",
	"IncrementBetweenViews":               "00191062",
	"DependantOnNoViewsProcessed":         "0019106A",
	"FieldOfViewInDetectorCells":          "0019106B",
	"ValueOfBackProjectionButton":         "00191070",
	"SetIfFatqEstimatesWereUsed":          "00191071",
	"ZChanAvgOverViews":                   "00191072",
	"AvgOfLeftRefChansOverViews":          "00191073",
	"MaxLeftChanOverViews":                "00191074",
	"AvgOfRightRefChansOverViews":         "00191075",
	"MaxRightChanOverViews":               "00191076",
	"SecondEcho":                          "0019107D",
	"NumberOfEchoes":                      "0019107E",
	"TableDelta":                          "0019107F",
	"Contiguous":                          "00191081",
	"PeakSAR":                             "00191084",
	"MonitorSAR":                          "00191085",
	"CardiacRepetitionTime":               "00191087",
	"ImagesPerCardiacCycle":               "00191088",
	"ActualReceiveGainAnalog":             "0019108A",
	"ActualReceiveGainDigital":            "0019108B",
	"DelayAfterTrigger":                   "0019108D",
	"Swappf":                              "0019108F",
	"PauseInterval":                       "00191090",
	"PulseTime":                           "00191091",
	"SliceOffsetOnFreqAxis":               "00191092",
	"CenterFrequency":                     "00191093",
	"TransmitGain":                        "00191094",
	"AnalogReceiverGain":                  "00191095",
	"DigitalReceiverGain":                 "00191096",
	"BitmapDefiningCVs":                   "00191097",
	"CenterFreqMethod":                    "00191098",
	"PulseSeqMode":                        "0019109B",
	"PulseSeqName":                        "0019109C",
	"PulseSeqDate":                        "0019109D",
	"InternalPulseSeqName":                "0019109E",
	"TransmittingCoil":                    "0019109F",
	"SurfaceCoilType":                     "001910A0",
	"ExtremityCoilFlag":                   "001910A1",
	"RawDataRunNumber":                    "001910A2",
	"CalibratedFieldStrength":             "001910A3",
	"SATFatWaterBone":                     "001910A4",
	"ReceiveBandwidth":                    "001910A5",
	"UserData01":                          "001910A7",
	"UserData02":                          "001910A8",
	"UserData03":                          "001910A9",
	"UserData04":                          "001910AA",
	"UserData05":                          "001910AB",
	"UserData06":                          "001910AC",
	"UserData07":                          "001910AD",
	"UserData08":                          "001910AE",
	"UserData09":                          "001910AF",
	"UserData10":                          "001910B0",
	"UserData11":                          "001910B1",
	"UserData12":                          "001910B2",
	"UserData13":                          "001910B3",
	"UserData14":                          "001910B4",
	"UserData15":                          "001910B5",
	"UserData16":                          "001910B6",
	"UserData17":                          "001910B7",
	"UserData18":                          "001910B8",
	"UserData19":                          "001910B9",
	"UserData20":                          "001910BA",
	"UserData21":                          "001910BB",
	"UserData22":                          "001910BC",
	"UserData23":                          "001910BD",
	"ProjectionAngle":                     "001910BE",
	"SaturationPlanes":                    "001910C0",
	"SurfaceCoilIntensity":                "001910C1",
	"SATLocationR":                        "001910C2",
	"SATLocationL":                        "001910C3",
	"SATLocationA":                        "001910C4",
	"SATLocationP":                        "001910C5",
	"SATLocationH":                        "001910C6",
	"SATLocationF":                        "001910C7",
	"SATThicknessRL":                      "001910C8",
	"SATThicknessAP":                      "001910C9",
	"SATThicknessHF":                      "001910CA",
	"PrescribedFlowAxis":                  "001910CB",
	"VelocityEncoding":                    "001910CC",
	"ThicknessDisclaimer":                 "001910CD",
	"PrescanType":                         "001910CE",
	"PrescanStatus":                       "001910CF",
	"RawDataType":                         "001910D0",
	"ProjectionAlgorithm":                 "001910D2",
	"FractionalEcho":                      "001910D5",
	"PrepPulse":                           "001910D6",
	"CardiacPhases":                       "001910D7",
	"VariableEchoflag":                    "001910D8",
	"ConcatenatedSAT":                     "001910D9",
	"ReferenceChannelUsed":                "001910DA",
	"BackProjectorCoefficient":            "001910DB",
	"PrimarySpeedCorrectionUsed":          "001910DC",
	"OverrangeCorrectionUsed":             "001910DD",
	"DynamicZAlphaValue":                  "001910DE",
	"UserData":                            "001910DF",
	"VelocityEncodeScale":                 "001910E2",
	"FastPhases":                          "001910F2",
	"TransmissionGain":                    "001910F9",
	"RelationshipGroupLength":             "00200000",
	"StudyInstanceUID":                    "0020000D",
	"SeriesInstanceUID":                   "0020000E",
	"StudyID":                             "00200010",
	"SeriesNumber":                        "00200011",
	"AcquisitionNumber":                   "00200012",
	"InstanceNumber":                      "00200013",
	"IsotopeNumber":                       "00200014",
	"PhaseNumber":                         "00200015",
	"IntervalNumber":                      "00200016",
	"TimeSlotNumber":                      "00200017",
	"AngleNumber":                         "00200018",
	"ItemNumber":                          "00200019",
	"PatientOrientation":                  "00200020",
	"OverlayNumber":                       "00200022",
	"CurveNumber":                         "00200024",
	"LookupTableNumber":                   "00200026",
	"ImagePosition":                       "00200030",
	"ImagePositionPatient":                "00200032",
	"ImageOrientation":                    "00200035",
	"ImageOrientationPatient":             "00200037",
	"Location":                            "00200050",
	"FrameOfReferenceUID":                 "00200052",
	"Laterality":                          "00200060",
	"ImageLaterality":                     "00200062",
	"ImageGeometryType":                   "00200070",
	"MaskingImage":                        "00200080",
	"TemporalPositionIdentifier":          "00200100",
	"NumberOfTemporalPositions":           "00200105",
	"TemporalResolution":                  "00200110",
	"SynchronizationFrameOfReferenceUID":  "00200200",
	"SeriesInStudy":                       "00201000",
	"AcquisitionsInSeries":                "00201001",
	"ImagesInAcquisition":                 "00201002",
	"ImagesInSeries":                      "00201003",
	"AcquisitionsInStudy":                 "00201004",
	"ImagesInStudy":                       "00201005",
	"Reference":                           "00201020",
	"PositionReferenceIndicator":          "00201040",
	"SliceLocation":                       "00201041",
	"OtherStudyNumbers":                   "00201070",
	"NumberOfPatientRelatedStudies":       "00201200",
	"NumberOfPatientRelatedSeries":        "00201202",
	"NumberOfPatientRelatedInstances":     "00201204",
	"NumberOfStudyRelatedSeries":          "00201206",
	"NumberOfStudyRelatedInstances":       "00201208",
	"NumberOfSeriesRelatedInstances":      "00201209",
	"SourceImageIDs":                      "002031xx",
	"ModifyingDeviceID":                   "00203401",
	"ModifiedImageID":                     "00203402",
	"ModifiedImageDate":                   "00203403",
	"ModifyingDeviceManufacturer":         "00203404",
	"ModifiedImageTime":                   "00203405",
	"ModifiedImageDescription":            "00203406",
	"ImageComments":                       "00204000",
	"OriginalImageIdentification":         "00205000",
	"OriginalImageIdentNomenclature":      "00205002",
	"StackID":                             "00209056",
	"InStackPositionNumber":               "00209057",
	"FrameAnatomySequence":                "00209071",
	"FrameLaterality":                     "00209072",
	"FrameContentSequence":                "00209111",
	"PlanePositionSequence":               "00209113",
	"PlaneOrientationSequence":            "00209116",
	"TemporalPositionIndex":               "00209128",
	"TriggerDelayTime":                    "00209153",
	"FrameAcquisitionNumber":              "00209156",
	"DimensionIndexValues":                "00209157",
	"FrameComments":                       "00209158",
	"ConcatenationUID":                    "00209161",
	"InConcatenationNumber":               "00209162",
	"InConcatenationTotalNumber":          "00209163",
	"DimensionOrganizationUID":            "00209164",
	"DimensionIndexPointer":               "00209165",
	"FunctionalGroupPointer":              "00209167",
	"DimensionIndexPrivateCreator":        "00209213",
	"DimensionOrganizationSequence":       "00209221",
	"DimensionIndexSequence":              "00209222",
	"ConcatenationFrameOffsetNumber":      "00209228",
	"FunctionalGroupPrivateCreator":       "00209238",
	"NominalPercentageOfCardiacPhase":     "00209241",
	"NominalPercentOfRespiratoryPhase":    "00209245",
	"StartingRespiratoryAmplitude":        "00209246",
	"StartingRespiratoryPhase":            "00209247",
	"EndingRespiratoryAmplitude":          "00209248",
	"EndingRespiratoryPhase":              "00209249",
	"RespiratoryTriggerType":              "00209250",
	"RRIntervalTimeNominal":               "00209251",
	"ActualCardiacTriggerDelayTime":       "00209252",
	"RespiratorySynchronizationSequence":  "00209253",
	"RespiratoryIntervalTime":             "00209254",
	"NominalRespiratoryTriggerDelayTime":  "00209255",
	"RespiratoryTriggerDelayThreshold":    "00209256",
	"ActualRespiratoryTriggerDelayTime":   "00209257",
	"ImagePositionVolume":                 "00209301",
	"ImageOrientationVolume":              "00209302",
	"ApexPosition":                        "00209308",
	"DimensionDescriptionLabel":           "00209421",
	"PatientOrientationInFrameSequence":   "00209450",
	"FrameLabel":                          "00209453",
	"AcquisitionIndex":                    "00209518",
	"ContributingSOPInstancesRefSeq":      "00209529",
	"ReconstructionIndex":                 "00209536",
	"SeriesFromWhichPrescribed":           "00211003",
	"SeriesRecordChecksum":                "00211007",
	"AcqreconRecordChecksum":              "00211019",
	"TableStartLocation":                  "00211020",
	"ImageFromWhichPrescribed":            "00211036",
	"ScreenFormat":                        "00211037",
	"AnatomicalReferenceForScout":         "0021104A",
	"LocationsInAcquisition":              "0021104F",
	"GraphicallyPrescribed":               "00211050",
	"RotationFromSourceXRot":              "00211051",
	"RotationFromSourceYRot":              "00211052",
	"RotationFromSourceZRot":              "00211053",
	"IntegerSlop":                         "00211056",
	"FloatSlop":                           "0021105F",
	"AutoWindowLevelAlpha":                "00211081",
	"AutoWindowLevelBeta":                 "00211082",
	"AutoWindowLevelWindow":               "00211083",
	"ToWindowLevelLevel":                  "00211084",
	"TubeFocalSpotPosition":               "00211090",
	"BiopsyPosition":                      "00211091",
	"BiopsyTLocation":                     "00211092",
	"BiopsyRefLocation":                   "00211093",
	"LightPathFilterPassThroughWavelen":   "00220001",
	"LightPathFilterPassBand":             "00220002",
	"ImagePathFilterPassThroughWavelen":   "00220003",
	"ImagePathFilterPassBand":             "00220004",
	"PatientEyeMovementCommanded":         "00220005",
	"PatientEyeMovementCommandCodeSeq":    "00220006",
	"SphericalLensPower":                  "00220007",
	"CylinderLensPower":                   "00220008",
	"CylinderAxis":                        "00220009",
	"EmmetropicMagnification":             "0022000A",
	"IntraOcularPressure":                 "0022000B",
	"HorizontalFieldOfView":               "0022000C",
	"PupilDilated":                        "0022000D",
	"DegreeOfDilation":                    "0022000E",
	"StereoBaselineAngle":                 "00220010",
	"StereoBaselineDisplacement":          "00220011",
	"StereoHorizontalPixelOffset":         "00220012",
	"StereoVerticalPixelOffset":           "00220013",
	"StereoRotation":                      "00220014",
	"AcquisitionDeviceTypeCodeSequence":   "00220015",
	"IlluminationTypeCodeSequence":        "00220016",
	"LightPathFilterTypeStackCodeSeq":     "00220017",
	"ImagePathFilterTypeStackCodeSeq":     "00220018",
	"LensesCodeSequence":                  "00220019",
	"ChannelDescriptionCodeSequence":      "0022001A",
	"RefractiveStateSequence":             "0022001B",
	"MydriaticAgentCodeSequence":          "0022001C",
	"RelativeImagePositionCodeSequence":   "0022001D",
	"StereoPairsSequence":                 "00220020",
	"LeftImageSequence":                   "00220021",
	"RightImageSequence":                  "00220022",
	"AxialLengthOfTheEye":                 "00220030",
	"OphthalmicFrameLocationSequence":     "00220031",
	"ReferenceCoordinates":                "00220032",
	"DepthSpatialResolution":              "00220035",
	"MaximumDepthDistortion":              "00220036",
	"AlongScanSpatialResolution":          "00220037",
	"MaximumAlongScanDistortion":          "00220038",
	"OphthalmicImageOrientation":          "00220039",
	"DepthOfTransverseImage":              "00220041",
	"MydriaticAgentConcUnitsSeq":          "00220042",
	"AcrossScanSpatialResolution":         "00220048",
	"MaximumAcrossScanDistortion":         "00220049",
	"MydriaticAgentConcentration":         "0022004E",
	"IlluminationWaveLength":              "00220055",
	"IlluminationPower":                   "00220056",
	"IlluminationBandwidth":               "00220057",
	"MydriaticAgentSequence":              "00220058",
	"NumberOfSeriesInStudy":               "00231001",
	"NumberOfUnarchivedSeries":            "00231002",
	"ReferenceImageField":                 "00231010",
	"SummaryImage":                        "00231050",
	"StartTimeSecsInFirstAxial":           "00231070",
	"NoofUpdatesToHeader":                 "00231074",
	"IndicatesIfTheStudyHasCompleteInfo":  "0023107D",
	"LastPulseSequenceUsed":               "00251006",
	"LandmarkCounter":                     "00251010",
	"NumberOfAcquisitions":                "00251011",
	"IndicatesNoofUpdatesToHeader":        "00251014",
	"SeriesCompleteFlag":                  "00251017",
	"NumberOfImagesArchived":              "00251018",
	"LastImageNumberUsed":                 "00251019",
	"PrimaryReceiverSuiteAndHost":         "0025101A",
	"ImageArchiveFlag":                    "00271006",
	"ScoutType":                           "00271010",
	"VmaMamp":                             "0027101C",
	"VmaPhase":                            "0027101D",
	"VmaMod":                              "0027101E",
	"VmaClip":                             "0027101F",
	"SmartScanOnOffFlag":                  "00271020",
	"ForeignImageRevision":                "00271030",
	"ImagingMode":                         "00271031",
	"PulseSequence":                       "00271032",
	"ImagingOptions":                      "00271033",
	"PlaneType":                           "00271035",
	"ObliquePlane":                        "00271036",
	"RASLetterOfImageLocation":            "00271040",
	"ImageLocation":                       "00271041",
	"CenterRCoordOfPlaneImage":            "00271042",
	"CenterACoordOfPlaneImage":            "00271043",
	"CenterSCoordOfPlaneImage":            "00271044",
	"NormalRCoord":                        "00271045",
	"NormalACoord":                        "00271046",
	"NormalSCoord":                        "00271047",
	"RCoordOfTopRightCorner":              "00271048",
	"ACoordOfTopRightCorner":              "00271049",
	"SCoordOfTopRightCorner":              "0027104A",
	"RCoordOfBottomRightCorner":           "0027104B",
	"ACoordOfBottomRightCorner":           "0027104C",
	"SCoordOfBottomRightCorner":           "0027104D",
	"TableEndLocation":                    "00271051",
	"RASLetterForSideOfImage":             "00271052",
	"RASLetterForAnteriorPosterior":       "00271053",
	"RASLetterForScoutStartLoc":           "00271054",
	"RASLetterForScoutEndLoc":             "00271055",
	"ImageDimensionX":                     "00271060",
	"ImageDimensionY":                     "00271061",
	"NumberOfExcitations":                 "00271062",
	"ImagePresentationGroupLength":        "00280000",
	"SamplesPerPixel":                     "00280002",
	"SamplesPerPixelUsed":                 "00280003",
	"PhotometricInterpretation":           "00280004",
	"ImageDimensions":                     "00280005",
	"PlanarConfiguration":                 "00280006",
	"NumberOfFrames":                      "00280008",
	"FrameIncrementPointer":               "00280009",
	"FrameDimensionPointer":               "0028000A",
	"Rows":                                "00280010",
	"Columns":                             "00280011",
	"Planes":                              "00280012",
	"UltrasoundColorDataPresent":          "00280014",
	"PixelSpacing":                        "00280030",
	"ZoomFactor":                          "00280031",
	"ZoomCenter":                          "00280032",
	"PixelAspectRatio":                    "00280034",
	"ImageFormat":                         "00280040",
	"ManipulatedImage":                    "00280050",
	"CorrectedImage":                      "00280051",
	"CompressionRecognitionCode":          "0028005F",
	"CompressionCode":                     "00280060",
	"CompressionOriginator":               "00280061",
	"CompressionLabel":                    "00280062",
	"CompressionDescription":              "00280063",
	"CompressionSequence":                 "00280065",
	"CompressionStepPointers":             "00280066",
	"RepeatInterval":                      "00280068",
	"BitsGrouped":                         "00280069",
	"PerimeterTable":                      "00280070",
	"PerimeterValue":                      "00280071",
	"PredictorRows":                       "00280080",
	"PredictorColumns":                    "00280081",
	"PredictorConstants":                  "00280082",
	"BlockedPixels":                       "00280090",
	"BlockRows":                           "00280091",
	"BlockColumns":                        "00280092",
	"RowOverlap":                          "00280093",
	"ColumnOverlap":                       "00280094",
	"BitsAllocated":                       "00280100",
	"BitsStored":                          "00280101",
	"HighBit":                             "00280102",
	"PixelRepresentation":                 "00280103",
	"SmallestValidPixelValue":             "00280104",
	"LargestValidPixelValue":              "00280105",
	"SmallestImagePixelValue":             "00280106",
	"LargestImagePixelValue":              "00280107",
	"SmallestPixelValueInSeries":          "00280108",
	"LargestPixelValueInSeries":           "00280109",
	"SmallestImagePixelValueInPlane":      "00280110",
	"LargestImagePixelValueInPlane":       "00280111",
	"PixelPaddingValue":                   "00280120",
	"PixelPaddingRangeLimit":              "00280121",
	"QualityControlImage":                 "00280300",
	"BurnedInAnnotation":                  "00280301",
	"TransformLabel":                      "00280400",
	"TransformVersionNumber":              "00280401",
	"NumberOfTransformSteps":              "00280402",
	"SequenceOfCompressedData":            "00280403",
	"DetailsOfCoefficients":               "00280404",
	"CoefficientCoding":                   "002804x2",
	"CoefficientCodingPointers":           "002804x3",
	"DCTLabel":                            "00280700",
	"DataBlockDescription":                "00280701",
	"DataBlock":                           "00280702",
	"NormalizationFactorFormat":           "00280710",
	"ZonalMapNumberFormat":                "00280720",
	"ZonalMapLocation":                    "00280721",
	"ZonalMapFormat":                      "00280722",
	"AdaptiveMapFormat":                   "00280730",
	"CodeNumberFormat":                    "00280740",
	"CodeLabel":                           "002808x0",
	"NumberOfTables":                      "002808x2",
	"CodeTableLocation":                   "002808x3",
	"BitsForCodeWord":                     "002808x4",
	"ImageDataLocation":                   "002808x8",
	"PixelSpacingCalibrationType":         "00280A02",
	"PixelSpacingCalibrationDescription":  "00280A04",
	"PixelIntensityRelationship":          "00281040",
	"PixelIntensityRelationshipSign":      "00281041",
	"WindowCenter":                        "00281050",
	"WindowWidth":                         "00281051",
	"RescaleIntercept":                    "00281052",
	"RescaleSlope":                        "00281053",
	"RescaleType":                         "00281054",
	"WindowCenterAndWidthExplanation":     "00281055",
	"VOILUTFunction":                      "00281056",
	"GrayScale":                           "00281080",
	"RecommendedViewingMode":              "00281090",
	"GrayLookupTableDescriptor":           "00281100",
	"RedPaletteColorTableDescriptor":      "00281101",
	"GreenPaletteColorTableDescriptor":    "00281102",
	"BluePaletteColorTableDescriptor":     "00281103",
	"LargeRedPaletteColorTableDescr":      "00281111",
	"LargeGreenPaletteColorTableDescr":    "00281112",
	"LargeBluePaletteColorTableDescr":     "00281113",
	"PaletteColorTableUID":                "00281199",
	"GrayLookupTableData":                 "00281200",
	"RedPaletteColorTableData":            "00281201",
	"GreenPaletteColorTableData":          "00281202",
	"BluePaletteColorTableData":           "00281203",
	"LargeRedPaletteColorTableData":       "00281211",
	"LargeGreenPaletteColorTableData":     "00281212",
	"LargeBluePaletteColorTableData":      "00281213",
	"LargePaletteColorLookupTableUID":     "00281214",
	"SegmentedRedColorTableData":          "00281221",
	"SegmentedGreenColorTableData":        "00281222",
	"SegmentedBlueColorTableData":         "00281223",
	"BreastImplantPresent":                "00281300",
	"PartialView":                         "00281350",
	"PartialViewDescription":              "00281351",
	"PartialViewCodeSequence":             "00281352",
	"SpatialLocationsPreserved":           "0028135A",
	"DataPathAssignment":                  "00281402",
	"BlendingLUT1Sequence":                "00281404",
	"BlendingWeightConstant":              "00281406",
	"BlendingLookupTableData":             "00281408",
	"BlendingLUT2Sequence":                "0028140C",
	"DataPathID":                          "0028140E",
	"RGBLUTTransferFunction":              "0028140F",
	"AlphaLUTTransferFunction":            "00281410",
	"ICCProfile":                          "00282000",
	"LossyImageCompressionRatio":          "00282112",
	"LossyImageCompressionMethod":         "00282114",
	"ModalityLUTSequence":                 "00283000",
	"LUTDescriptor":                       "00283002",
	"LUTExplanation":                      "00283003",
	"ModalityLUTType":                     "00283004",
	"LUTData":                             "00283006",
	"VOILUTSequence":                      "00283010",
	"SoftcopyVOILUTSequence":              "00283110",
	"ImagePresentationComments":           "00284000",
	"BiPlaneAcquisitionSequence":          "00285000",
	"RepresentativeFrameNumber":           "00286010",
	"FrameNumbersOfInterest":              "00286020",
	"FrameOfInterestDescription":          "00286022",
	"FrameOfInterestType":                 "00286023",
	"MaskPointers":                        "00286030",
	"RWavePointer":                        "00286040",
	"MaskSubtractionSequence":             "00286100",
	"MaskOperation":                       "00286101",
	"ApplicableFrameRange":                "00286102",
	"MaskFrameNumbers":                    "00286110",
	"ContrastFrameAveraging":              "00286112",
	"MaskSubPixelShift":                   "00286114",
	"TIDOffset":                           "00286120",
	"MaskOperationExplanation":            "00286190",
	"PixelDataProviderURL":                "00287FE0",
	"DataPointRows":                       "00289001",
	"DataPointColumns":                    "00289002",
	"SignalDomainColumns":                 "00289003",
	"LargestMonochromePixelValue":         "00289099",
	"DataRepresentation":                  "00289108",
	"PixelMeasuresSequence":               "00289110",
	"FrameVOILUTSequence":                 "00289132",
	"PixelValueTransformationSequence":    "00289145",
	"SignalDomainRows":                    "00289235",
	"DisplayFilterPercentage":             "00289411",
	"FramePixelShiftSequence":             "00289415",
	"SubtractionItemID":                   "00289416",
	"PixelIntensityRelationshipLUTSeq":    "00289422",
	"FramePixelDataPropertiesSequence":    "00289443",
	"GeometricalProperties":               "00289444",
	"GeometricMaximumDistortion":          "00289445",
	"ImageProcessingApplied":              "00289446",
	"MaskSelectionMode":                   "00289454",
	"LUTFunction":                         "00289474",
	"MaskVisibilityPercentage":            "00289478",
	"PixelShiftSequence":                  "00289501",
	"RegionPixelShiftSequence":            "00289502",
	"VerticesOfTheRegion":                 "00289503",
	"PixelShiftFrameRange":                "00289506",
	"LUTFrameRange":                       "00289507",
	"ImageToEquipmentMappingMatrix":       "00289520",
	"EquipmentCoordinateSystemID":         "00289537",
	"LowerRangeOfPixels1a":                "00291004",
	"LowerRangeOfPixels1b":                "00291005",
	"LowerRangeOfPixels1c":                "00291006",
	"LowerRangeOfPixels1d":                "00291007",
	"LowerRangeOfPixels1e":                "00291008",
	"LowerRangeOfPixels1f":                "00291009",
	"LowerRangeOfPixels1g":                "0029100A",
	"LowerRangeOfPixels1h":                "00291015",
	"LowerRangeOfPixels1i":                "00291016",
	"LowerRangeOfPixels2":                 "00291017",
	"UpperRangeOfPixels2":                 "00291018",
	"LenOfTotHdrInBytes":                  "0029101A",
	"VersionOfTheHdrStruct":               "00291026",
	"AdvantageCompOverflow":               "00291034",
	"AdvantageCompUnderflow":              "00291035",
	"StudyGroupLength":                    "00320000",
	"StudyStatusID":                       "0032000A",
	"StudyPriorityID":                     "0032000C",
	"StudyIDIssuer":                       "00320012",
	"StudyVerifiedDate":                   "00320032",
	"StudyVerifiedTime":                   "00320033",
	"StudyReadDate":                       "00320034",
	"StudyReadTime":                       "00320035",
	"ScheduledStudyStartDate":             "00321000",
	"ScheduledStudyStartTime":             "00321001",
	"ScheduledStudyStopDate":              "00321010",
	"ScheduledStudyStopTime":              "00321011",
	"ScheduledStudyLocation":              "00321020",
	"ScheduledStudyLocationAETitle":       "00321021",
	"ReasonForStudy":                      "00321030",
	"RequestingPhysicianIDSequence":       "00321031",
	"RequestingPhysician":                 "00321032",
	"RequestingService":                   "00321033",
	"StudyArrivalDate":                    "00321040",
	"StudyArrivalTime":                    "00321041",
	"StudyCompletionDate":                 "00321050",
	"StudyCompletionTime":                 "00321051",
	"StudyComponentStatusID":              "00321055",
	"RequestedProcedureDescription":       "00321060",
	"RequestedProcedureCodeSequence":      "00321064",
	"RequestedContrastAgent":              "00321070",
	"StudyComments":                       "00324000",
	"ReferencedPatientAliasSequence":      "00380004",
	"VisitStatusID":                       "00380008",
	"AdmissionID":                         "00380010",
	"IssuerOfAdmissionID":                 "00380011",
	"RouteOfAdmissions":                   "00380016",
	"ScheduledAdmissionDate":              "0038001A",
	"ScheduledAdmissionTime":              "0038001B",
	"ScheduledDischargeDate":              "0038001C",
	"ScheduledDischargeTime":              "0038001D",
	"ScheduledPatientInstitResidence":     "0038001E",
	"AdmittingDate":                       "00380020",
	"AdmittingTime":                       "00380021",
	"DischargeDate":                       "00380030",
	"DischargeTime":                       "00380032",
	"DischargeDiagnosisDescription":       "00380040",
	"DischargeDiagnosisCodeSequence":      "00380044",
	"SpecialNeeds":                        "00380050",
	"ServiceEpisodeID":                    "00380060",
	"IssuerOfServiceEpisodeID":            "00380061",
	"ServiceEpisodeDescription":           "00380062",
	"PertinentDocumentsSequence":          "00380100",
	"CurrentPatientLocation":              "00380300",
	"PatientInstitutionResidence":         "00380400",
	"PatientState":                        "00380500",
	"PatientClinicalTrialParticipSeq":     "00380502",
	"VisitComments":                       "00384000",
	"WaveformOriginality":                 "003A0004",
	"NumberOfWaveformChannels":            "003A0005",
	"NumberOfWaveformSamples":             "003A0010",
	"SamplingFrequency":                   "003A001A",
	"MultiplexGroupLabel":                 "003A0020",
	"ChannelDefinitionSequence":           "003A0200",
	"WaveformChannelNumber":               "003A0202",
	"ChannelLabel":                        "003A0203",
	"ChannelStatus":                       "003A0205",
	"ChannelSourceSequence":               "003A0208",
	"ChannelSourceModifiersSequence":      "003A0209",
	"SourceWaveformSequence":              "003A020A",
	"ChannelDerivationDescription":        "003A020C",
	"ChannelSensitivity":                  "003A0210",
	"ChannelSensitivityUnitsSequence":     "003A0211",
	"ChannelSensitivityCorrectionFactor":  "003A0212",
	"ChannelBaseline":                     "003A0213",
	"ChannelTimeSkew":                     "003A0214",
	"ChannelSampleSkew":                   "003A0215",
	"ChannelOffset":                       "003A0218",
	"WaveformBitsStored":                  "003A021A",
	"FilterLowFrequency":                  "003A0220",
	"FilterHighFrequency":                 "003A0221",
	"NotchFilterFrequency":                "003A0222",
	"NotchFilterBandwidth":                "003A0223",
	"WaveformDataDisplayScale":            "003A0230",
	"WaveformDisplayBkgCIELabValue":       "003A0231",
	"WaveformPresentationGroupSequence":   "003A0240",
	"PresentationGroupNumber":             "003A0241",
	"ChannelDisplaySequence":              "003A0242",
	"ChannelRecommendDisplayCIELabValue":  "003A0244",
	"ChannelPosition":                     "003A0245",
	"DisplayShadingFlag":                  "003A0246",
	"FractionalChannelDisplayScale":       "003A0247",
	"AbsoluteChannelDisplayScale":         "003A0248",
	"MultiplexAudioChannelsDescrCodeSeq":  "003A0300",
	"ChannelIdentificationCode":           "003A0301",
	"ChannelMode":                         "003A0302",
	"ScheduledStationAETitle":             "00400001",
	"ScheduledProcedureStepStartDate":     "00400002",
	"ScheduledProcedureStepStartTime":     "00400003",
	"ScheduledProcedureStepEndDate":       "00400004",
	"ScheduledProcedureStepEndTime":       "00400005",
	"ScheduledPerformingPhysiciansName":   "00400006",
	"ScheduledProcedureStepDescription":   "00400007",
	"ScheduledProtocolCodeSequence":       "00400008",
	"ScheduledProcedureStepID":            "00400009",
	"StageCodeSequence":                   "0040000A",
	"ScheduledPerformingPhysicianIDSeq":   "0040000B",
	"ScheduledStationName":                "00400010",
	"ScheduledProcedureStepLocation":      "00400011",
	"PreMedication":                       "00400012",
	"ScheduledProcedureStepStatus":        "00400020",
	"LocalNamespaceEntityID":              "00400031",
	"UniversalEntityID":                   "00400032",
	"UniversalEntityIDType":               "00400033",
	"IdentifierTypeCode":                  "00400035",
	"AssigningFacilitySequence":           "00400036",
	"ScheduledProcedureStepSequence":      "00400100",
	"ReferencedNonImageCompositeSOPSeq":   "00400220",
	"PerformedStationAETitle":             "00400241",
	"PerformedStationName":                "00400242",
	"PerformedLocation":                   "00400243",
	"PerformedProcedureStepStartDate":     "00400244",
	"PerformedProcedureStepStartTime":     "00400245",
	"PerformedProcedureStepEndDate":       "00400250",
	"PerformedProcedureStepEndTime":       "00400251",
	"PerformedProcedureStepStatus":        "00400252",
	"PerformedProcedureStepID":            "00400253",
	"PerformedProcedureStepDescription":   "00400254",
	"PerformedProcedureTypeDescription":   "00400255",
	"PerformedProtocolCodeSequence":       "00400260",
	"PerformedProtocolType":               "00400261",
	"ScheduledStepAttributesSequence":     "00400270",
	"RequestAttributesSequence":           "00400275",
	"CommentsOnPerformedProcedureStep":    "00400280",
	"ProcStepDiscontinueReasonCodeSeq":    "00400281",
	"QuantitySequence":                    "00400293",
	"Quantity":                            "00400294",
	"MeasuringUnitsSequence":              "00400295",
	"BillingItemSequence":                 "00400296",
	"TotalTimeOfFluoroscopy":              "00400300",
	"TotalNumberOfExposures":              "00400301",
	"EntranceDose":                        "00400302",
	"ExposedArea":                         "00400303",
	"DistanceSourceToEntrance":            "00400306",
	"DistanceSourceToSupport":             "00400307",
	"ExposureDoseSequence":                "0040030E",
	"CommentsOnRadiationDose":             "00400310",
	"XRayOutput":                          "00400312",
	"HalfValueLayer":                      "00400314",
	"OrganDose":                           "00400316",
	"OrganExposed":                        "00400318",
	"BillingProcedureStepSequence":        "00400320",
	"FilmConsumptionSequence":             "00400321",
	"BillingSuppliesAndDevicesSequence":   "00400324",
	"PerformedSeriesSequence":             "00400340",
	"CommentsOnScheduledProcedureStep":    "00400400",
	"ProtocolContextSequence":             "00400440",
	"ContentItemModifierSequence":         "00400441",
	"SpecimenAccessionNumber":             "0040050A",
	"ContainerIdentifier":                 "00400512",
	"ContainerDescription":                "0040051A",
	"SpecimenSequence":                    "00400550",
	"SpecimenIdentifier":                  "00400551",
	"SpecimenDescriptionSequenceTrial":    "00400552",
	"SpecimenDescriptionTrial":            "00400553",
	"SpecimenUID":                         "00400554",
	"AcquisitionContextSequence":          "00400555",
	"AcquisitionContextDescription":       "00400556",
	"SpecimenTypeCodeSequence":            "0040059A",
	"SpecimenShortDescription":            "00400600",
	"SlideIdentifier":                     "004006FA",
	"ImageCenterPointCoordinatesSeq":      "0040071A",
	"XOffsetInSlideCoordinateSystem":      "0040072A",
	"YOffsetInSlideCoordinateSystem":      "0040073A",
	"ZOffsetInSlideCoordinateSystem":      "0040074A",
	"PixelSpacingSequence":                "004008D8",
	"CoordinateSystemAxisCodeSequence":    "004008DA",
	"MeasurementUnitsCodeSequence":        "004008EA",
	"VitalStainCodeSequenceTrial":         "004009F8",
	"RequestedProcedureID":                "00401001",
	"ReasonForRequestedProcedure":         "00401002",
	"RequestedProcedurePriority":          "00401003",
	"PatientTransportArrangements":        "00401004",
	"RequestedProcedureLocation":          "00401005",
	"PlacerOrderNumberProcedure":          "00401006",
	"FillerOrderNumberProcedure":          "00401007",
	"ConfidentialityCode":                 "00401008",
	"ReportingPriority":                   "00401009",
	"ReasonForRequestedProcedureCodeSeq":  "0040100A",
	"NamesOfIntendedRecipientsOfResults":  "00401010",
	"IntendedRecipientsOfResultsIDSeq":    "00401011",
	"PersonIdentificationCodeSequence":    "00401101",
	"PersonAddress":                       "00401102",
	"PersonTelephoneNumbers":              "00401103",
	"RequestedProcedureComments":          "00401400",
	"ReasonForImagingServiceRequest":      "00402001",
	"IssueDateOfImagingServiceRequest":    "00402004",
	"IssueTimeOfImagingServiceRequest":    "00402005",
	"PlacerOrderNumImagingServiceReq":     "00402006",
	"FillerOrderNumImagingServiceReq":     "00402007",
	"OrderEnteredBy":                      "00402008",
	"OrderEntererLocation":                "00402009",
	"OrderCallbackPhoneNumber":            "00402010",
	"ImagingServiceRequestComments":       "00402400",
	"ConfidentialityOnPatientDataDescr":   "00403001",
	"GenPurposeScheduledProcStepStatus":   "00404001",
	"GenPurposePerformedProcStepStatus":   "00404002",
	"GenPurposeSchedProcStepPriority":     "00404003",
	"SchedProcessingApplicationsCodeSeq":  "00404004",
	"SchedProcedureStepStartDateAndTime":  "00404005",
	"MultipleCopiesFlag":                  "00404006",
	"PerformedProcessingAppsCodeSeq":      "00404007",
	"HumanPerformerCodeSequence":          "00404009",
	"SchedProcStepModificationDateTime":   "00404010",
	"ExpectedCompletionDateAndTime":       "00404011",
	"ResultingGenPurposePerfProcStepSeq":  "00404015",
	"RefGenPurposeSchedProcStepSeq":       "00404016",
	"ScheduledWorkitemCodeSequence":       "00404018",
	"PerformedWorkitemCodeSequence":       "00404019",
	"InputAvailabilityFlag":               "00404020",
	"InputInformationSequence":            "00404021",
	"RelevantInformationSequence":         "00404022",
	"RefGenPurSchedProcStepTransUID":      "00404023",
	"ScheduledStationNameCodeSequence":    "00404025",
	"ScheduledStationClassCodeSequence":   "00404026",
	"SchedStationGeographicLocCodeSeq":    "00404027",
	"PerformedStationNameCodeSequence":    "00404028",
	"PerformedStationClassCodeSequence":   "00404029",
	"PerformedStationGeogLocCodeSeq":      "00404030",
	"RequestedSubsequentWorkItemCodeSeq":  "00404031",
	"NonDICOMOutputCodeSequence":          "00404032",
	"OutputInformationSequence":           "00404033",
	"ScheduledHumanPerformersSequence":    "00404034",
	"ActualHumanPerformersSequence":       "00404035",
	"HumanPerformersOrganization":         "00404036",
	"HumanPerformerName":                  "00404037",
	"RawDataHandling":                     "00404040",
	"EntranceDoseInMilliGy":               "00408302",
	"RefImageRealWorldValueMappingSeq":    "00409094",
	"RealWorldValueMappingSequence":       "00409096",
	"PixelValueMappingCodeSequence":       "00409098",
	"LUTLabel":                            "00409210",
	"RealWorldValueLastValueMapped":       "00409211",
	"RealWorldValueLUTData":               "00409212",
	"RealWorldValueFirstValueMapped":      "00409216",
	"RealWorldValueIntercept":             "00409224",
	"RealWorldValueSlope":                 "00409225",
	"RelationshipType":                    "0040A010",
	"VerifyingOrganization":               "0040A027",
	"VerificationDateTime":                "0040A030",
	"ObservationDateTime":                 "0040A032",
	"ValueType":                           "0040A040",
	"ConceptNameCodeSequence":             "0040A043",
	"ContinuityOfContent":                 "0040A050",
	"VerifyingObserverSequence":           "0040A073",
	"VerifyingObserverName":               "0040A075",
	"AuthorObserverSequence":              "0040A078",
	"ParticipantSequence":                 "0040A07A",
	"CustodialOrganizationSequence":       "0040A07C",
	"ParticipationType":                   "0040A080",
	"ParticipationDateTime":               "0040A082",
	"ObserverType":                        "0040A084",
	"VerifyingObserverIdentCodeSequence":  "0040A088",
	"EquivalentCDADocumentSequence":       "0040A090",
	"ReferencedWaveformChannels":          "0040A0B0",
	"DateTime":                            "0040A120",
	"Date":                                "0040A121",
	"Time":                                "0040A122",
	"PersonName":                          "0040A123",
	"UID":                                 "0040A124",
	"TemporalRangeType":                   "0040A130",
	"ReferencedSamplePositions":           "0040A132",
	"ReferencedFrameNumbers":              "0040A136",
	"ReferencedTimeOffsets":               "0040A138",
	"ReferencedDateTime":                  "0040A13A",
	"TextValue":                           "0040A160",
	"ConceptCodeSequence":                 "0040A168",
	"PurposeOfReferenceCodeSequence":      "0040A170",
	"AnnotationGroupNumber":               "0040A180",
	"ModifierCodeSequence":                "0040A195",
	"MeasuredValueSequence":               "0040A300",
	"NumericValueQualifierCodeSequence":   "0040A301",
	"NumericValue":                        "0040A30A",
	"AddressTrial":                        "0040A353",
	"TelephoneNumberTrial":                "0040A354",
	"PredecessorDocumentsSequence":        "0040A360",
	"ReferencedRequestSequence":           "0040A370",
	"PerformedProcedureCodeSequence":      "0040A372",
	"CurrentRequestedProcEvidenceSeq":     "0040A375",
	"PertinentOtherEvidenceSequence":      "0040A385",
	"HL7StructuredDocumentRefSeq":         "0040A390",
	"CompletionFlag":                      "0040A491",
	"CompletionFlagDescription":           "0040A492",
	"VerificationFlag":                    "0040A493",
	"ArchiveRequested":                    "0040A494",
	"PreliminaryFlag":                     "0040A496",
	"ContentTemplateSequence":             "0040A504",
	"IdenticalDocumentsSequence":          "0040A525",
	"ContentSequence":                     "0040A730",
	"AnnotationSequence":                  "0040B020",
	"TemplateIdentifier":                  "0040DB00",
	"TemplateVersion":                     "0040DB06",
	"TemplateLocalVersion":                "0040DB07",
	"TemplateExtensionFlag":               "0040DB0B",
	"TemplateExtensionOrganizationUID":    "0040DB0C",
	"TemplateExtensionCreatorUID":         "0040DB0D",
	"ReferencedContentItemIdentifier":     "0040DB73",
	"HL7InstanceIdentifier":               "0040E001",
	"HL7DocumentEffectiveTime":            "0040E004",
	"HL7DocumentTypeCodeSequence":         "0040E006",
	"RetrieveURI":                         "0040E010",
	"RetrieveLocationUID":                 "0040E011",
	"DocumentTitle":                       "00420010",
	"EncapsulatedDocument":                "00420011",
	"MIMETypeOfEncapsulatedDocument":      "00420012",
	"SourceInstanceSequence":              "00420013",
	"ListOfMIMETypes":                     "00420014",
	"BitmapOfPrescanOptions":              "00431001",
	"GradientOffsetInX":                   "00431002",
	"GradientOffsetInY":                   "00431003",
	"GradientOffsetInZ":                   "00431004",
	"ImgIsOriginalOrUnoriginal":           "00431005",
	"NumberOfEPIShots":                    "00431006",
	"ViewsPerSegment":                     "00431007",
	"RespiratoryRateBpm":                  "00431008",
	"RespiratoryTriggerPoint":             "00431009",
	"TypeOfReceiverUsed":                  "0043100A",
	"PeakRateOfChangeOfGradientField":     "0043100B",
	"LimitsInUnitsOfPercent":              "0043100C",
	"PSDEstimatedLimit":                   "0043100D",
	"PSDEstimatedLimitInTeslaPerSecond":   "0043100E",
	"Saravghead":                          "0043100F",
	"WindowValue":                         "00431010",
	"TotalInputViews":                     "00431011",
	"XRayChain":                           "00431012",
	"DeconKernelParameters":               "00431013",
	"CalibrationParameters":               "00431014",
	"TotalOutputViews":                    "00431015",
	"NumberOfOverranges":                  "00431016",
	"IBHImageScaleFactors":                "00431017",
	"BBHCoefficients":                     "00431018",
	"NumberOfBBHChainsToBlend":            "00431019",
	"StartingChannelNumber":               "0043101A",
	"PpscanParameters":                    "0043101B",
	"GEImageIntegrity":                    "0043101C",
	"LevelValue":                          "0043101D",
	"DeltaStartTime":                      "0043101E",
	"MaxOverrangesInAView":                "0043101F",
	"AvgOverrangesAllViews":               "00431020",
	"CorrectedAfterGlowTerms":             "00431021",
	"ReferenceChannels":                   "00431025",
	"NoViewsRefChansBlocked":              "00431026",
	"ScanPitchRatio":                      "00431027",
	"UniqueImageIden":                     "00431028",
	"HistogramTables":                     "00431029",
	"UserDefinedData":                     "0043102A",
	"PrivateScanOptions":                  "0043102B",
	"EffectiveEchoSpacing":                "0043102C",
	"StringSlopField1":                    "0043102D",
	"StringSlopField2":                    "0043102E",
	"RACordOfTargetReconCenter":           "00431031",
	"NegScanspacing":                      "00431033",
	"OffsetFrequency":                     "00431034",
	"UserUsageTag":                        "00431035",
	"UserFillMapMSW":                      "00431036",
	"UserFillMapLSW":                      "00431037",
	"User2548":                            "00431038",
	"SlopInt69":                           "00431039",
	"TriggerOnPosition":                   "00431040",
	"DegreeOfRotation":                    "00431041",
	"DASTriggerSource":                    "00431042",
	"DASFpaGain":                          "00431043",
	"DASOutputSource":                     "00431044",
	"DASAdInput":                          "00431045",
	"DASCalMode":                          "00431046",
	"DASCalFrequency":                     "00431047",
	"DASRegXm":                            "00431048",
	"DASAutoZero":                         "00431049",
	"StartingChannelOfView":               "0043104A",
	"DASXmPattern":                        "0043104B",
	"TGGCTriggerMode":                     "0043104C",
	"StartScanToXrayOnDelay":              "0043104D",
	"DurationOfXrayOn":                    "0043104E",
	"SlopInt1017":                         "00431060",
	"ScannerStudyEntityUID":               "00431061",
	"ScannerStudyID":                      "00431062",
	"ScannerTableEntry":                   "0043106f",
	"ProductPackageIdentifier":            "00440001",
	"SubstanceAdministrationApproval":     "00440002",
	"ApprovalStatusFurtherDescription":    "00440003",
	"ApprovalStatusDateTime":              "00440004",
	"ProductTypeCodeSequence":             "00440007",
	"ProductName":                         "00440008",
	"ProductDescription":                  "00440009",
	"ProductLotIdentifier":                "0044000A",
	"ProductExpirationDateTime":           "0044000B",
	"SubstanceAdministrationDateTime":     "00440010",
	"SubstanceAdministrationNotes":        "00440011",
	"SubstanceAdministrationDeviceID":     "00440012",
	"ProductParameterSequence":            "00440013",
	"SubstanceAdminParameterSeq":          "00440019",
	"NumberOfMacroRowsInDetector":         "00451001",
	"MacroWidthAtISOCenter":               "00451002",
	"DASType":                             "00451003",
	"DASGain":                             "00451004",
	"DASTemperature":                      "00451005",
	"TableDirectionInOrOut":               "00451006",
	"ZSmoothingFactor":                    "00451007",
	"ViewWeightingMode":                   "00451008",
	"SigmaRowNumberWhichRowsWereUsed":     "00451009",
	"MinimumDasValueFoundInTheScanData":   "0045100A",
	"MaximumOffsetShiftValueUsed":         "0045100B",
	"NumberOfViewsShifted":                "0045100C",
	"ZTrackingFlag":                       "0045100D",
	"MeanZError":                          "0045100E",
	"ZTrackingMaximumError":               "0045100F",
	"StartingViewForRow2a":                "00451010",
	"NumberOfViewsInRow2a":                "00451011",
	"StartingViewForRow1a":                "00451012",
	"SigmaMode":                           "00451013",
	"NumberOfViewsInRow1a":                "00451014",
	"StartingViewForRow2b":                "00451015",
	"NumberOfViewsInRow2b":                "00451016",
	"StartingViewForRow1b":                "00451017",
	"NumberOfViewsInRow1b":                "00451018",
	"AirFilterCalibrationDate":            "00451019",
	"AirFilterCalibrationTime":            "0045101A",
	"PhantomCalibrationDate":              "0045101B",
	"PhantomCalibrationTime":              "0045101C",
	"ZSlopeCalibrationDate":               "0045101D",
	"ZSlopeCalibrationTime":               "0045101E",
	"CrosstalkCalibrationDate":            "0045101F",
	"CrosstalkCalibrationTime":            "00451020",
	"IterboneOptionFlag":                  "00451021",
	"PeristalticFlagOption":               "00451022",
	"LensDescription":                     "00460012",
	"RightLensSequence":                   "00460014",
	"LeftLensSequence":                    "00460015",
	"CylinderSequence":                    "00460018",
	"PrismSequence":                       "00460028",
	"HorizontalPrismPower":                "00460030",
	"HorizontalPrismBase":                 "00460032",
	"VerticalPrismPower":                  "00460034",
	"VerticalPrismBase":                   "00460036",
	"LensSegmentType":                     "00460038",
	"OpticalTransmittance":                "00460040",
	"ChannelWidth":                        "00460042",
	"PupilSize":                           "00460044",
	"CornealSize":                         "00460046",
	"DistancePupillaryDistance":           "00460060",
	"NearPupillaryDistance":               "00460062",
	"OtherPupillaryDistance":              "00460064",
	"RadiusOfCurvature":                   "00460075",
	"KeratometricPower":                   "00460076",
	"KeratometricAxis":                    "00460077",
	"BackgroundColor":                     "00460092",
	"Optotype":                            "00460094",
	"OptotypePresentation":                "00460095",
	"AddNearSequence":                     "00460100",
	"AddIntermediateSequence":             "00460101",
	"AddOtherSequence":                    "00460102",
	"AddPower":                            "00460104",
	"ViewingDistance":                     "00460106",
	"ViewingDistanceType":                 "00460125",
	"VisualAcuityModifiers":               "00460135",
	"DecimalVisualAcuity":                 "00460137",
	"OptotypeDetailedDefinition":          "00460139",
	"SpherePower":                         "00460146",
	"CylinderPower":                       "00460147",
	"CalibrationImage":                    "00500004",
	"DeviceSequence":                      "00500010",
	"DeviceLength":                        "00500014",
	"ContainerComponentWidth":             "00500015",
	"DeviceDiameter":                      "00500016",
	"DeviceDiameterUnits":                 "00500017",
	"DeviceVolume":                        "00500018",
	"InterMarkerDistance":                 "00500019",
	"ContainerComponentID":                "0050001B",
	"DeviceDescription":                   "00500020",
	"EnergyWindowVector":                  "00540010",
	"NumberOfEnergyWindows":               "00540011",
	"EnergyWindowInformationSequence":     "00540012",
	"EnergyWindowRangeSequence":           "00540013",
	"EnergyWindowLowerLimit":              "00540014",
	"EnergyWindowUpperLimit":              "00540015",
	"RadiopharmaceuticalInformationSeq":   "00540016",
	"ResidualSyringeCounts":               "00540017",
	"EnergyWindowName":                    "00540018",
	"DetectorVector":                      "00540020",
	"NumberOfDetectors":                   "00540021",
	"DetectorInformationSequence":         "00540022",
	"PhaseVector":                         "00540030",
	"NumberOfPhases":                      "00540031",
	"PhaseInformationSequence":            "00540032",
	"NumberOfFramesInPhase":               "00540033",
	"PhaseDelay":                          "00540036",
	"PauseBetweenFrames":                  "00540038",
	"PhaseDescription":                    "00540039",
	"RotationVector":                      "00540050",
	"NumberOfRotations":                   "00540051",
	"RotationInformationSequence":         "00540052",
	"NumberOfFramesInRotation":            "00540053",
	"RRIntervalVector":                    "00540060",
	"NumberOfRRIntervals":                 "00540061",
	"GatedInformationSequence":            "00540062",
	"DataInformationSequence":             "00540063",
	"TimeSlotVector":                      "00540070",
	"NumberOfTimeSlots":                   "00540071",
	"TimeSlotInformationSequence":         "00540072",
	"TimeSlotTime":                        "00540073",
	"SliceVector":                         "00540080",
	"NumberOfSlices":                      "00540081",
	"AngularViewVector":                   "00540090",
	"TimeSliceVector":                     "00540100",
	"NumberOfTimeSlices":                  "00540101",
	"StartAngle":                          "00540200",
	"TypeOfDetectorMotion":                "00540202",
	"TriggerVector":                       "00540210",
	"NumberOfTriggersInPhase":             "00540211",
	"ViewCodeSequence":                    "00540220",
	"ViewModifierCodeSequence":            "00540222",
	"RadionuclideCodeSequence":            "00540300",
	"AdministrationRouteCodeSequence":     "00540302",
	"RadiopharmaceuticalCodeSequence":     "00540304",
	"CalibrationDataSequence":             "00540306",
	"EnergyWindowNumber":                  "00540308",
	"ImageID":                             "00540400",
	"PatientOrientationCodeSequence":      "00540410",
	"PatientOrientationModifierCodeSeq":   "00540412",
	"PatientGantryRelationshipCodeSeq":    "00540414",
	"SliceProgressionDirection":           "00540500",
	"SeriesType":                          "00541000",
	"Units":                               "00541001",
	"CountsSource":                        "00541002",
	"ReprojectionMethod":                  "00541004",
	"RandomsCorrectionMethod":             "00541100",
	"AttenuationCorrectionMethod":         "00541101",
	"DecayCorrection":                     "00541102",
	"ReconstructionMethod":                "00541103",
	"DetectorLinesOfResponseUsed":         "00541104",
	"ScatterCorrectionMethod":             "00541105",
	"AxialAcceptance":                     "00541200",
	"AxialMash":                           "00541201",
	"TransverseMash":                      "00541202",
	"DetectorElementSize":                 "00541203",
	"CoincidenceWindowWidth":              "00541210",
	"SecondaryCountsType":                 "00541220",
	"FrameReferenceTime":                  "00541300",
	"PrimaryCountsAccumulated":            "00541310",
	"SecondaryCountsAccumulated":          "00541311",
	"SliceSensitivityFactor":              "00541320",
	"DecayFactor":                         "00541321",
	"DoseCalibrationFactor":               "00541322",
	"ScatterFractionFactor":               "00541323",
	"DeadTimeFactor":                      "00541324",
	"ImageIndex":                          "00541330",
	"CountsIncluded":                      "00541400",
	"DeadTimeCorrectionFlag":              "00541401",
	"HistogramSequence":                   "00603000",
	"HistogramNumberOfBins":               "00603002",
	"HistogramFirstBinValue":              "00603004",
	"HistogramLastBinValue":               "00603006",
	"HistogramBinWidth":                   "00603008",
	"HistogramExplanation":                "00603010",
	"HistogramData":                       "00603020",
	"SegmentationType":                    "00620001",
	"SegmentSequence":                     "00620002",
	"SegmentedPropertyCategoryCodeSeq":    "00620003",
	"SegmentLabel":                        "00620005",
	"SegmentDescription":                  "00620006",
	"SegmentAlgorithmType":                "00620008",
	"SegmentAlgorithmName":                "00620009",
	"SegmentIdentificationSequence":       "0062000A",
	"ReferencedSegmentNumber":             "0062000B",
	"RecommendedDisplayGrayscaleValue":    "0062000C",
	"RecommendedDisplayCIELabValue":       "0062000D",
	"MaximumFractionalValue":              "0062000E",
	"SegmentedPropertyTypeCodeSequence":   "0062000F",
	"SegmentationFractionalType":          "00620010",
	"DeformableRegistrationSequence":      "00640002",
	"SourceFrameOfReferenceUID":           "00640003",
	"DeformableRegistrationGridSequence":  "00640005",
	"GridDimensions":                      "00640007",
	"GridResolution":                      "00640008",
	"VectorGridData":                      "00640009",
	"PreDeformationMatrixRegistSeq":       "0064000F",
	"PostDeformationMatrixRegistSeq":      "00640010",
	"NumberOfSurfaces":                    "00660001",
	"SurfaceSequence":                     "00660002",
	"SurfaceNumber":                       "00660003",
	"SurfaceComments":                     "00660004",
	"SurfaceProcessing":                   "00660009",
	"SurfaceProcessingRatio":              "0066000A",
	"FiniteVolume":                        "0066000E",
	"Manifold":                            "00660010",
	"SurfacePointsSequence":               "00660011",
	"NumberOfSurfacePoints":               "00660015",
	"PointCoordinatesData":                "00660016",
	"PointPositionAccuracy":               "00660017",
	"MeanPointDistance":                   "00660018",
	"MaximumPointDistance":                "00660019",
	"AxisOfRotation":                      "0066001B",
	"CenterOfRotation":                    "0066001C",
	"NumberOfVectors":                     "0066001E",
	"VectorDimensionality":                "0066001F",
	"VectorAccuracy":                      "00660020",
	"VectorCoordinateData":                "00660021",
	"TrianglePointIndexList":              "00660023",
	"EdgePointIndexList":                  "00660024",
	"VertexPointIndexList":                "00660025",
	"TriangleStripSequence":               "00660026",
	"TriangleFanSequence":                 "00660027",
	"LineSequence":                        "00660028",
	"PrimitivePointIndexList":             "00660029",
	"SurfaceCount":                        "0066002A",
	"AlgorithmFamilyCodeSequ":             "0066002F",
	"AlgorithmVersion":                    "00660031",
	"AlgorithmParameters":                 "00660032",
	"FacetSequence":                       "00660034",
	"AlgorithmName":                       "00660036",
	"GraphicAnnotationSequence":           "00700001",
	"GraphicLayer":                        "00700002",
	"BoundingBoxAnnotationUnits":          "00700003",
	"AnchorPointAnnotationUnits":          "00700004",
	"GraphicAnnotationUnits":              "00700005",
	"UnformattedTextValue":                "00700006",
	"TextObjectSequence":                  "00700008",
	"GraphicObjectSequence":               "00700009",
	"BoundingBoxTopLeftHandCorner":        "00700010",
	"BoundingBoxBottomRightHandCorner":    "00700011",
	"BoundingBoxTextHorizJustification":   "00700012",
	"AnchorPoint":                         "00700014",
	"AnchorPointVisibility":               "00700015",
	"GraphicDimensions":                   "00700020",
	"NumberOfGraphicPoints":               "00700021",
	"GraphicData":                         "00700022",
	"GraphicType":                         "00700023",
	"GraphicFilled":                       "00700024",
	"ImageRotationRetired":                "00700040",
	"ImageHorizontalFlip":                 "00700041",
	"ImageRotation":                       "00700042",
	"DisplayedAreaTopLeftTrial":           "00700050",
	"DisplayedAreaBottomRightTrial":       "00700051",
	"DisplayedAreaTopLeft":                "00700052",
	"DisplayedAreaBottomRight":            "00700053",
	"DisplayedAreaSelectionSequence":      "0070005A",
	"GraphicLayerSequence":                "00700060",
	"GraphicLayerOrder":                   "00700062",
	"GraphicLayerRecDisplayGraysclValue":  "00700066",
	"GraphicLayerRecDisplayRGBValue":      "00700067",
	"GraphicLayerDescription":             "00700068",
	"ContentLabel":                        "00700080",
	"ContentDescription":                  "00700081",
	"PresentationCreationDate":            "00700082",
	"PresentationCreationTime":            "00700083",
	"ContentCreatorName":                  "00700084",
	"ContentCreatorIDCodeSequence":        "00700086",
	"PresentationSizeMode":                "00700100",
	"PresentationPixelSpacing":            "00700101",
	"PresentationPixelAspectRatio":        "00700102",
	"PresentationPixelMagRatio":           "00700103",
	"ShapeType":                           "00700306",
	"RegistrationSequence":                "00700308",
	"MatrixRegistrationSequence":          "00700309",
	"MatrixSequence":                      "0070030A",
	"FrameOfRefTransformationMatrixType":  "0070030C",
	"RegistrationTypeCodeSequence":        "0070030D",
	"FiducialDescription":                 "0070030F",
	"FiducialIdentifier":                  "00700310",
	"FiducialIdentifierCodeSequence":      "00700311",
	"ContourUncertaintyRadius":            "00700312",
	"UsedFiducialsSequence":               "00700314",
	"GraphicCoordinatesDataSequence":      "00700318",
	"FiducialUID":                         "0070031A",
	"FiducialSetSequence":                 "0070031C",
	"FiducialSequence":                    "0070031E",
	"GraphicLayerRecomDisplayCIELabVal":   "00700401",
	"BlendingSequence":                    "00700402",
	"RelativeOpacity":                     "00700403",
	"ReferencedSpatialRegistrationSeq":    "00700404",
	"BlendingPosition":                    "00700405",
	"HangingProtocolName":                 "00720002",
	"HangingProtocolDescription":          "00720004",
	"HangingProtocolLevel":                "00720006",
	"HangingProtocolCreator":              "00720008",
	"HangingProtocolCreationDateTime":     "0072000A",
	"HangingProtocolDefinitionSequence":   "0072000C",
	"HangingProtocolUserIDCodeSequence":   "0072000E",
	"HangingProtocolUserGroupName":        "00720010",
	"SourceHangingProtocolSequence":       "00720012",
	"NumberOfPriorsReferenced":            "00720014",
	"ImageSetsSequence":                   "00720020",
	"ImageSetSelectorSequence":            "00720022",
	"ImageSetSelectorUsageFlag":           "00720024",
	"SelectorAttribute":                   "00720026",
	"SelectorValueNumber":                 "00720028",
	"TimeBasedImageSetsSequence":          "00720030",
	"ImageSetNumber":                      "00720032",
	"ImageSetSelectorCategory":            "00720034",
	"RelativeTime":                        "00720038",
	"RelativeTimeUnits":                   "0072003A",
	"AbstractPriorValue":                  "0072003C",
	"AbstractPriorCodeSequence":           "0072003E",
	"ImageSetLabel":                       "00720040",
	"SelectorAttributeVR":                 "00720050",
	"SelectorSequencePointer":             "00720052",
	"SelectorSeqPointerPrivateCreator":    "00720054",
	"SelectorAttributePrivateCreator":     "00720056",
	"SelectorATValue":                     "00720060",
	"SelectorCSValue":                     "00720062",
	"SelectorISValue":                     "00720064",
	"SelectorLOValue":                     "00720066",
	"SelectorLTValue":                     "00720068",
	"SelectorPNValue":                     "0072006A",
	"SelectorSHValue":                     "0072006C",
	"SelectorSTValue":                     "0072006E",
	"SelectorUTValue":                     "00720070",
	"SelectorDSValue":                     "00720072",
	"SelectorFDValue":                     "00720074",
	"SelectorFLValue":                     "00720076",
	"SelectorULValue":                     "00720078",
	"SelectorUSValue":                     "0072007A",
	"SelectorSLValue":                     "0072007C",
	"SelectorSSValue":                     "0072007E",
	"SelectorCodeSequenceValue":           "00720080",
	"NumberOfScreens":                     "00720100",
	"NominalScreenDefinitionSequence":     "00720102",
	"NumberOfVerticalPixels":              "00720104",
	"NumberOfHorizontalPixels":            "00720106",
	"DisplayEnvironmentSpatialPosition":   "00720108",
	"ScreenMinimumGrayscaleBitDepth":      "0072010A",
	"ScreenMinimumColorBitDepth":          "0072010C",
	"ApplicationMaximumRepaintTime":       "0072010E",
	"DisplaySetsSequence":                 "00720200",
	"DisplaySetNumber":                    "00720202",
	"DisplaySetLabel":                     "00720203",
	"DisplaySetPresentationGroup":         "00720204",
	"DisplaySetPresentationGroupDescr":    "00720206",
	"PartialDataDisplayHandling":          "00720208",
	"SynchronizedScrollingSequence":       "00720210",
	"DisplaySetScrollingGroup":            "00720212",
	"NavigationIndicatorSequence":         "00720214",
	"NavigationDisplaySet":                "00720216",
	"ReferenceDisplaySets":                "00720218",
	"ImageBoxesSequence":                  "00720300",
	"ImageBoxNumber":                      "00720302",
	"ImageBoxLayoutType":                  "00720304",
	"ImageBoxTileHorizontalDimension":     "00720306",
	"ImageBoxTileVerticalDimension":       "00720308",
	"ImageBoxScrollDirection":             "00720310",
	"ImageBoxSmallScrollType":             "00720312",
	"ImageBoxSmallScrollAmount":           "00720314",
	"ImageBoxLargeScrollType":             "00720316",
	"ImageBoxLargeScrollAmount":           "00720318",
	"ImageBoxOverlapPriority":             "00720320",
	"CineRelativeToRealTime":              "00720330",
	"FilterOperationsSequence":            "00720400",
	"FilterByCategory":                    "00720402",
	"FilterByAttributePresence":           "00720404",
	"FilterByOperator":                    "00720406",
	"SynchronizedImageBoxList":            "00720432",
	"TypeOfSynchronization":               "00720434",
	"BlendingOperationType":               "00720500",
	"ReformattingOperationType":           "00720510",
	"ReformattingThickness":               "00720512",
	"ReformattingInterval":                "00720514",
	"ReformattingOpInitialViewDir":        "00720516",
	"RenderingType3D":                     "00720520",
	"SortingOperationsSequence":           "00720600",
	"SortByCategory":                      "00720602",
	"SortingDirection":                    "00720604",
	"DisplaySetPatientOrientation":        "00720700",
	"VOIType":                             "00720702",
	"PseudoColorType":                     "00720704",
	"ShowGrayscaleInverted":               "00720706",
	"ShowImageTrueSizeFlag":               "00720710",
	"ShowGraphicAnnotationFlag":           "00720712",
	"ShowPatientDemographicsFlag":         "00720714",
	"ShowAcquisitionTechniquesFlag":       "00720716",
	"DisplaySetHorizontalJustification":   "00720717",
	"DisplaySetVerticalJustification":     "00720718",
	"UnifiedProcedureStepState":           "00741000",
	"UPSProgressInformationSequence":      "00741002",
	"UnifiedProcedureStepProgress":        "00741004",
	"UnifiedProcedureStepProgressDescr":   "00741006",
	"UnifiedProcedureStepComURISeq":       "00741008",
	"ContactURI":                          "0074100a",
	"ContactDisplayName":                  "0074100c",
	"BeamTaskSequence":                    "00741020",
	"BeamTaskType":                        "00741022",
	"BeamOrderIndex":                      "00741024",
	"DeliveryVerificationImageSequence":   "00741030",
	"VerificationImageTiming":             "00741032",
	"DoubleExposureFlag":                  "00741034",
	"DoubleExposureOrdering":              "00741036",
	"DoubleExposureMeterset":              "00741038",
	"DoubleExposureFieldDelta":            "0074103A",
	"RelatedReferenceRTImageSequence":     "00741040",
	"GeneralMachineVerificationSequence":  "00741042",
	"ConventionalMachineVerificationSeq":  "00741044",
	"IonMachineVerificationSequence":      "00741046",
	"FailedAttributesSequence":            "00741048",
	"OverriddenAttributesSequence":        "0074104A",
	"ConventionalControlPointVerifySeq":   "0074104C",
	"IonControlPointVerificationSeq":      "0074104E",
	"AttributeOccurrenceSequence":         "00741050",
	"AttributeOccurrencePointer":          "00741052",
	"AttributeItemSelector":               "00741054",
	"AttributeOccurrencePrivateCreator":   "00741056",
	"ScheduledProcedureStepPriority":      "00741200",
	"WorklistLabel":                       "00741202",
	"ProcedureStepLabel":                  "00741204",
	"ScheduledProcessingParametersSeq":    "00741210",
	"PerformedProcessingParametersSeq":    "00741212",
	"UPSPerformedProcedureSequence":       "00741216",
	"RelatedProcedureStepSequence":        "00741220",
	"ProcedureStepRelationshipType":       "00741222",
	"DeletionLock":                        "00741230",
	"ReceivingAE":                         "00741234",
	"RequestingAE":                        "00741236",
	"ReasonForCancellation":               "00741238",
	"SCPStatus":                           "00741242",
	"SubscriptionListStatus":              "00741244",
	"UPSListStatus":                       "00741246",
	"StorageMediaFileSetID":               "00880130",
	"StorageMediaFileSetUID":              "00880140",
	"IconImageSequence":                   "00880200",
	"TopicTitle":                          "00880904",
	"TopicSubject":                        "00880906",
	"TopicAuthor":                         "00880910",
	"TopicKeywords":                       "00880912",
	"SOPInstanceStatus":                   "01000410",
	"SOPAuthorizationDateAndTime":         "01000420",
	"SOPAuthorizationComment":             "01000424",
	"AuthorizationEquipmentCertNumber":    "01000426",
	"MACIDNumber":                         "04000005",
	"MACCalculationTransferSyntaxUID":     "04000010",
	"MACAlgorithm":                        "04000015",
	"DataElementsSigned":                  "04000020",
	"DigitalSignatureUID":                 "04000100",
	"DigitalSignatureDateTime":            "04000105",
	"CertificateType":                     "04000110",
	"CertificateOfSigner":                 "04000115",
	"Signature":                           "04000120",
	"CertifiedTimestampType":              "04000305",
	"CertifiedTimestamp":                  "04000310",
	"DigitalSignaturePurposeCodeSeq":      "04000401",
	"ReferencedDigitalSignatureSeq":       "04000402",
	"ReferencedSOPInstanceMACSeq":         "04000403",
	"MAC":                                 "04000404",
	"EncryptedAttributesSequence":         "04000500",
	"EncryptedContentTransferSyntaxUID":   "04000510",
	"EncryptedContent":                    "04000520",
	"ModifiedAttributesSequence":          "04000550",
	"OriginalAttributesSequence":          "04000561",
	"AttributeModificationDateTime":       "04000562",
	"ModifyingSystem":                     "04000563",
	"SourceOfPreviousValues":              "04000564",
	"ReasonForTheAttributeModification":   "04000565",
	"EscapeTriplet":                       "1000xxx0",
	"RunLengthTriplet":                    "1000xxx1",
	"HuffmanTableSize":                    "1000xxx2",
	"HuffmanTableTriplet":                 "1000xxx3",
	"ShiftTableSize":                      "1000xxx4",
	"ShiftTableTriplet":                   "1000xxx5",
	"ZonalMap":                            "1010xxxx",
	"NumberOfCopies":                      "20000010",
	"PrinterConfigurationSequence":        "2000001E",
	"PrintPriority":                       "20000020",
	"MediumType":                          "20000030",
	"FilmDestination":                     "20000040",
	"FilmSessionLabel":                    "20000050",
	"MemoryAllocation":                    "20000060",
	"MaximumMemoryAllocation":             "20000061",
	"ColorImagePrintingFlag":              "20000062",
	"CollationFlag":                       "20000063",
	"AnnotationFlag":                      "20000065",
	"ImageOverlayFlag":                    "20000067",
	"PresentationLUTFlag":                 "20000069",
	"ImageBoxPresentationLUTFlag":         "2000006A",
	"MemoryBitDepth":                      "200000A0",
	"PrintingBitDepth":                    "200000A1",
	"MediaInstalledSequence":              "200000A2",
	"OtherMediaAvailableSequence":         "200000A4",
	"SupportedImageDisplayFormatSeq":      "200000A8",
	"ReferencedFilmBoxSequence":           "20000500",
	"ReferencedStoredPrintSequence":       "20000510",
	"ImageDisplayFormat":                  "20100010",
	"AnnotationDisplayFormatID":           "20100030",
	"FilmOrientation":                     "20100040",
	"FilmSizeID":                          "20100050",
	"PrinterResolutionID":                 "20100052",
	"DefaultPrinterResolutionID":          "20100054",
	"MagnificationType":                   "20100060",
	"SmoothingType":                       "20100080",
	"DefaultMagnificationType":            "201000A6",
	"OtherMagnificationTypesAvailable":    "201000A7",
	"DefaultSmoothingType":                "201000A8",
	"OtherSmoothingTypesAvailable":        "201000A9",
	"BorderDensity":                       "20100100",
	"EmptyImageDensity":                   "20100110",
	"MinDensity":                          "20100120",
	"MaxDensity":                          "20100130",
	"Trim":                                "20100140",
	"ConfigurationInformation":            "20100150",
	"ConfigurationInformationDescr":       "20100152",
	"MaximumCollatedFilms":                "20100154",
	"Illumination":                        "2010015E",
	"ReflectedAmbientLight":               "20100160",
	"PrinterPixelSpacing":                 "20100376",
	"ReferencedFilmSessionSequence":       "20100500",
	"ReferencedImageBoxSequence":          "20100510",
	"ReferencedBasicAnnotationBoxSeq":     "20100520",
	"ImageBoxPosition":                    "20200010",
	"Polarity":                            "20200020",
	"RequestedImageSize":                  "20200030",
	"RequestedDecimateCropBehavior":       "20200040",
	"RequestedResolutionID":               "20200050",
	"RequestedImageSizeFlag":              "202000A0",
	"DecimateCropResult":                  "202000A2",
	"BasicGrayscaleImageSequence":         "20200110",
	"BasicColorImageSequence":             "20200111",
	"ReferencedImageOverlayBoxSequence":   "20200130",
	"ReferencedVOILUTBoxSequence":         "20200140",
	"AnnotationPosition":                  "20300010",
	"TextString":                          "20300020",
	"ReferencedOverlayPlaneSequence":      "20400010",
	"ReferencedOverlayPlaneGroups":        "20400011",
	"OverlayPixelDataSequence":            "20400020",
	"OverlayMagnificationType":            "20400060",
	"OverlaySmoothingType":                "20400070",
	"OverlayOrImageMagnification":         "20400072",
	"MagnifyToNumberOfColumns":            "20400074",
	"OverlayForegroundDensity":            "20400080",
	"OverlayBackgroundDensity":            "20400082",
	"OverlayMode":                         "20400090",
	"ThresholdDensity":                    "20400100",
	"PresentationLUTSequence":             "20500010",
	"PresentationLUTShape":                "20500020",
	"ReferencedPresentationLUTSequence":   "20500500",
	"PrintJobID":                          "21000010",
	"ExecutionStatus":                     "21000020",
	"ExecutionStatusInfo":                 "21000030",
	"CreationDate":                        "21000040",
	"CreationTime":                        "21000050",
	"Originator":                          "21000070",
	"DestinationAE":                       "21000140",
	"OwnerID":                             "21000160",
	"NumberOfFilms":                       "21000170",
	"ReferencedPrintJobSequence":          "21000500",
	"PrinterStatus":                       "21100010",
	"PrinterStatusInfo":                   "21100020",
	"PrinterName":                         "21100030",
	"PrintQueueID":                        "21100099",
	"QueueStatus":                         "21200010",
	"PrintJobDescriptionSequence":         "21200050",
	"PrintManagementCapabilitiesSeq":      "21300010",
	"PrinterCharacteristicsSequence":      "21300015",
	"FilmBoxContentSequence":              "21300030",
	"ImageBoxContentSequence":             "21300040",
	"AnnotationContentSequence":           "21300050",
	"ImageOverlayBoxContentSequence":      "21300060",
	"PresentationLUTContentSequence":      "21300080",
	"ProposedStudySequence":               "213000A0",
	"OriginalImageSequence":               "213000C0",
	"LabelFromInfoExtractedFromInstance":  "22000001",
	"LabelText":                           "22000002",
	"LabelStyleSelection":                 "22000003",
	"MediaDisposition":                    "22000004",
	"BarcodeValue":                        "22000005",
	"BarcodeSymbology":                    "22000006",
	"AllowMediaSplitting":                 "22000007",
	"IncludeNonDICOMObjects":              "22000008",
	"IncludeDisplayApplication":           "22000009",
	"SaveCompInstancesAfterMediaCreate":   "2200000A",
	"TotalNumberMediaPiecesCreated":       "2200000B",
	"RequestedMediaApplicationProfile":    "2200000C",
	"ReferencedStorageMediaSequence":      "2200000D",
	"FailureAttributes":                   "2200000E",
	"AllowLossyCompression":               "2200000F",
	"RequestPriority":                     "22000020",
	"RTImageLabel":                        "30020002",
	"RTImageName":                         "30020003",
	"RTImageDescription":                  "30020004",
	"ReportedValuesOrigin":                "3002000A",
	"RTImagePlane":                        "3002000C",
	"XRayImageReceptorTranslation":        "3002000D",
	"XRayImageReceptorAngle":              "3002000E",
	"RTImageOrientation":                  "30020010",
	"ImagePlanePixelSpacing":              "30020011",
	"RTImagePosition":                     "30020012",
	"RadiationMachineName":                "30020020",
	"RadiationMachineSAD":                 "30020022",
	"RadiationMachineSSD":                 "30020024",
	"RTImageSID":                          "30020026",
	"SourceToReferenceObjectDistance":     "30020028",
	"FractionNumber":                      "30020029",
	"ExposureSequence":                    "30020030",
	"MetersetExposure":                    "30020032",
	"DiaphragmPosition":                   "30020034",
	"FluenceMapSequence":                  "30020040",
	"FluenceDataSource":                   "30020041",
	"FluenceDataScale":                    "30020042",
	"FluenceMode":                         "30020051",
	"FluenceModeID":                       "30020052",
	"DVHType":                             "30040001",
	"DoseUnits":                           "30040002",
	"DoseType":                            "30040004",
	"DoseComment":                         "30040006",
	"NormalizationPoint":                  "30040008",
	"DoseSummationType":                   "3004000A",
	"GridFrameOffsetVector":               "3004000C",
	"DoseGridScaling":                     "3004000E",
	"RTDoseROISequence":                   "30040010",
	"DoseValue":                           "30040012",
	"TissueHeterogeneityCorrection":       "30040014",
	"DVHNormalizationPoint":               "30040040",
	"DVHNormalizationDoseValue":           "30040042",
	"DVHSequence":                         "30040050",
	"DVHDoseScaling":                      "30040052",
	"DVHVolumeUnits":                      "30040054",
	"DVHNumberOfBins":                     "30040056",
	"DVHData":                             "30040058",
	"DVHReferencedROISequence":            "30040060",
	"DVHROIContributionType":              "30040062",
	"DVHMinimumDose":                      "30040070",
	"DVHMaximumDose":                      "30040072",
	"DVHMeanDose":                         "30040074",
	"StructureSetLabel":                   "30060002",
	"StructureSetName":                    "30060004",
	"StructureSetDescription":             "30060006",
	"StructureSetDate":                    "30060008",
	"StructureSetTime":                    "30060009",
	"ReferencedFrameOfReferenceSequence":  "30060010",
	"RTReferencedStudySequence":           "30060012",
	"RTReferencedSeriesSequence":          "30060014",
	"ContourImageSequence":                "30060016",
	"StructureSetROISequence":             "30060020",
	"ROINumber":                           "30060022",
	"ReferencedFrameOfReferenceUID":       "30060024",
	"ROIName":                             "30060026",
	"ROIDescription":                      "30060028",
	"ROIDisplayColor":                     "3006002A",
	"ROIVolume":                           "3006002C",
	"RTRelatedROISequence":                "30060030",
	"RTROIRelationship":                   "30060033",
	"ROIGenerationAlgorithm":              "30060036",
	"ROIGenerationDescription":            "30060038",
	"ROIContourSequence":                  "30060039",
	"ContourSequence":                     "30060040",
	"ContourGeometricType":                "30060042",
	"ContourSlabThickness":                "30060044",
	"ContourOffsetVector":                 "30060045",
	"NumberOfContourPoints":               "30060046",
	"ContourNumber":                       "30060048",
	"AttachedContours":                    "30060049",
	"ContourData":                         "30060050",
	"RTROIObservationsSequence":           "30060080",
	"ObservationNumber":                   "30060082",
	"ReferencedROINumber":                 "30060084",
	"ROIObservationLabel":                 "30060085",
	"RTROIIdentificationCodeSequence":     "30060086",
	"ROIObservationDescription":           "30060088",
	"RelatedRTROIObservationsSequence":    "300600A0",
	"RTROIInterpretedType":                "300600A4",
	"ROIInterpreter":                      "300600A6",
	"ROIPhysicalPropertiesSequence":       "300600B0",
	"ROIPhysicalProperty":                 "300600B2",
	"ROIPhysicalPropertyValue":            "300600B4",
	"ROIElementalCompositionSequence":     "300600B6",
	"ROIElementalCompAtomicNumber":        "300600B7",
	"ROIElementalCompAtomicMassFraction":  "300600B8",
	"FrameOfReferenceRelationshipSeq":     "300600C0",
	"RelatedFrameOfReferenceUID":          "300600C2",
	"FrameOfReferenceTransformType":       "300600C4",
	"FrameOfReferenceTransformMatrix":     "300600C6",
	"FrameOfReferenceTransformComment":    "300600C8",
	"MeasuredDoseReferenceSequence":       "30080010",
	"MeasuredDoseDescription":             "30080012",
	"MeasuredDoseType":                    "30080014",
	"MeasuredDoseValue":                   "30080016",
	"TreatmentSessionBeamSequence":        "30080020",
	"TreatmentSessionIonBeamSequence":     "30080021",
	"CurrentFractionNumber":               "30080022",
	"TreatmentControlPointDate":           "30080024",
	"TreatmentControlPointTime":           "30080025",
	"TreatmentTerminationStatus":          "3008002A",
	"TreatmentTerminationCode":            "3008002B",
	"TreatmentVerificationStatus":         "3008002C",
	"ReferencedTreatmentRecordSequence":   "30080030",
	"SpecifiedPrimaryMeterset":            "30080032",
	"SpecifiedSecondaryMeterset":          "30080033",
	"DeliveredPrimaryMeterset":            "30080036",
	"DeliveredSecondaryMeterset":          "30080037",
	"SpecifiedTreatmentTime":              "3008003A",
	"DeliveredTreatmentTime":              "3008003B",
	"ControlPointDeliverySequence":        "30080040",
	"IonControlPointDeliverySequence":     "30080041",
	"SpecifiedMeterset":                   "30080042",
	"DeliveredMeterset":                   "30080044",
	"MetersetRateSet":                     "30080045",
	"MetersetRateDelivered":               "30080046",
	"ScanSpotMetersetsDelivered":          "30080047",
	"DoseRateDelivered":                   "30080048",
	"TreatmentSummaryCalcDoseRefSeq":      "30080050",
	"CumulativeDoseToDoseReference":       "30080052",
	"FirstTreatmentDate":                  "30080054",
	"MostRecentTreatmentDate":             "30080056",
	"NumberOfFractionsDelivered":          "3008005A",
	"OverrideSequence":                    "30080060",
	"ParameterSequencePointer":            "30080061",
	"OverrideParameterPointer":            "30080062",
	"ParameterItemIndex":                  "30080063",
	"MeasuredDoseReferenceNumber":         "30080064",
	"ParameterPointer":                    "30080065",
	"OverrideReason":                      "30080066",
	"CorrectedParameterSequence":          "30080068",
	"CorrectionValue":                     "3008006A",
	"CalculatedDoseReferenceSequence":     "30080070",
	"CalculatedDoseReferenceNumber":       "30080072",
	"CalculatedDoseReferenceDescription":  "30080074",
	"CalculatedDoseReferenceDoseValue":    "30080076",
	"StartMeterset":                       "30080078",
	"EndMeterset":                         "3008007A",
	"ReferencedMeasuredDoseReferenceSeq":  "30080080",
	"ReferencedMeasuredDoseReferenceNum":  "30080082",
	"ReferencedCalculatedDoseRefSeq":      "30080090",
	"ReferencedCalculatedDoseRefNumber":   "30080092",
	"BeamLimitingDeviceLeafPairsSeq":      "300800A0",
	"RecordedWedgeSequence":               "300800B0",
	"RecordedCompensatorSequence":         "300800C0",
	"RecordedBlockSequence":               "300800D0",
	"TreatmentSummaryMeasuredDoseRefSeq":  "300800E0",
	"RecordedSnoutSequence":               "300800F0",
	"RecordedRangeShifterSequence":        "300800F2",
	"RecordedLateralSpreadingDeviceSeq":   "300800F4",
	"RecordedRangeModulatorSequence":      "300800F6",
	"RecordedSourceSequence":              "30080100",
	"SourceSerialNumber":                  "30080105",
	"TreatmentSessionAppSetupSeq":         "30080110",
	"ApplicationSetupCheck":               "30080116",
	"RecordedBrachyAccessoryDeviceSeq":    "30080120",
	"ReferencedBrachyAccessoryDeviceNum":  "30080122",
	"RecordedChannelSequence":             "30080130",
	"SpecifiedChannelTotalTime":           "30080132",
	"DeliveredChannelTotalTime":           "30080134",
	"SpecifiedNumberOfPulses":             "30080136",
	"DeliveredNumberOfPulses":             "30080138",
	"SpecifiedPulseRepetitionInterval":    "3008013A",
	"DeliveredPulseRepetitionInterval":    "3008013C",
	"RecordedSourceApplicatorSequence":    "30080140",
	"ReferencedSourceApplicatorNumber":    "30080142",
	"RecordedChannelShieldSequence":       "30080150",
	"ReferencedChannelShieldNumber":       "30080152",
	"BrachyControlPointDeliveredSeq":      "30080160",
	"SafePositionExitDate":                "30080162",
	"SafePositionExitTime":                "30080164",
	"SafePositionReturnDate":              "30080166",
	"SafePositionReturnTime":              "30080168",
	"CurrentTreatmentStatus":              "30080200",
	"TreatmentStatusComment":              "30080202",
	"FractionGroupSummarySequence":        "30080220",
	"ReferencedFractionNumber":            "30080223",
	"FractionGroupType":                   "30080224",
	"BeamStopperPosition":                 "30080230",
	"FractionStatusSummarySequence":       "30080240",
	"TreatmentDate":                       "30080250",
	"TreatmentTime":                       "30080251",
	"RTPlanLabel":                         "300A0002",
	"RTPlanName":                          "300A0003",
	"RTPlanDescription":                   "300A0004",
	"RTPlanDate":                          "300A0006",
	"RTPlanTime":                          "300A0007",
	"TreatmentProtocols":                  "300A0009",
	"PlanIntent":                          "300A000A",
	"TreatmentSites":                      "300A000B",
	"RTPlanGeometry":                      "300A000C",
	"PrescriptionDescription":             "300A000E",
	"DoseReferenceSequence":               "300A0010",
	"DoseReferenceNumber":                 "300A0012",
	"DoseReferenceUID":                    "300A0013",
	"DoseReferenceStructureType":          "300A0014",
	"NominalBeamEnergyUnit":               "300A0015",
	"DoseReferenceDescription":            "300A0016",
	"DoseReferencePointCoordinates":       "300A0018",
	"NominalPriorDose":                    "300A001A",
	"DoseReferenceType":                   "300A0020",
	"ConstraintWeight":                    "300A0021",
	"DeliveryWarningDose":                 "300A0022",
	"DeliveryMaximumDose":                 "300A0023",
	"TargetMinimumDose":                   "300A0025",
	"TargetPrescriptionDose":              "300A0026",
	"TargetMaximumDose":                   "300A0027",
	"TargetUnderdoseVolumeFraction":       "300A0028",
	"OrganAtRiskFullVolumeDose":           "300A002A",
	"OrganAtRiskLimitDose":                "300A002B",
	"OrganAtRiskMaximumDose":              "300A002C",
	"OrganAtRiskOverdoseVolumeFraction":   "300A002D",
	"ToleranceTableSequence":              "300A0040",
	"ToleranceTableNumber":                "300A0042",
	"ToleranceTableLabel":                 "300A0043",
	"GantryAngleTolerance":                "300A0044",
	"BeamLimitingDeviceAngleTolerance":    "300A0046",
	"BeamLimitingDeviceToleranceSeq":      "300A0048",
	"BeamLimitingDevicePositionTol":       "300A004A",
	"SnoutPositionTolerance":              "300A004B",
	"PatientSupportAngleTolerance":        "300A004C",
	"TableTopEccentricAngleTolerance":     "300A004E",
	"TableTopPitchAngleTolerance":         "300A004F",
	"TableTopRollAngleTolerance":          "300A0050",
	"TableTopVerticalPositionTolerance":   "300A0051",
	"TableTopLongitudinalPositionTol":     "300A0052",
	"TableTopLateralPositionTolerance":    "300A0053",
	"RTPlanRelationship":                  "300A0055",
	"FractionGroupSequence":               "300A0070",
	"FractionGroupNumber":                 "300A0071",
	"FractionGroupDescription":            "300A0072",
	"NumberOfFractionsPlanned":            "300A0078",
	"NumberFractionPatternDigitsPerDay":   "300A0079",
	"RepeatFractionCycleLength":           "300A007A",
	"FractionPattern":                     "300A007B",
	"NumberOfBeams":                       "300A0080",
	"BeamDoseSpecificationPoint":          "300A0082",
	"BeamDose":                            "300A0084",
	"BeamMeterset":                        "300A0086",
	"BeamDosePointDepth":                  "300A0088",
	"BeamDosePointEquivalentDepth":        "300A0089",
	"BeamDosePointSSD":                    "300A008A",
	"NumberOfBrachyApplicationSetups":     "300A00A0",
	"BrachyAppSetupDoseSpecPoint":         "300A00A2",
	"BrachyApplicationSetupDose":          "300A00A4",
	"BeamSequence":                        "300A00B0",
	"TreatmentMachineName":                "300A00B2",
	"PrimaryDosimeterUnit":                "300A00B3",
	"SourceAxisDistance":                  "300A00B4",
	"BeamLimitingDeviceSequence":          "300A00B6",
	"RTBeamLimitingDeviceType":            "300A00B8",
	"SourceToBeamLimitingDeviceDistance":  "300A00BA",
	"IsocenterToBeamLimitingDeviceDist":   "300A00BB",
	"NumberOfLeafJawPairs":                "300A00BC",
	"LeafPositionBoundaries":              "300A00BE",
	"BeamNumber":                          "300A00C0",
	"BeamName":                            "300A00C2",
	"BeamDescription":                     "300A00C3",
	"BeamType":                            "300A00C4",
	"RadiationType":                       "300A00C6",
	"HighDoseTechniqueType":               "300A00C7",
	"ReferenceImageNumber":                "300A00C8",
	"PlannedVerificationImageSequence":    "300A00CA",
	"ImagingDeviceSpecificAcqParams":      "300A00CC",
	"TreatmentDeliveryType":               "300A00CE",
	"NumberOfWedges":                      "300A00D0",
	"WedgeSequence":                       "300A00D1",
	"WedgeNumber":                         "300A00D2",
	"WedgeType":                           "300A00D3",
	"WedgeID":                             "300A00D4",
	"WedgeAngle":                          "300A00D5",
	"WedgeFactor":                         "300A00D6",
	"TotalWedgeTrayWaterEquivThickness":   "300A00D7",
	"WedgeOrientation":                    "300A00D8",
	"IsocenterToWedgeTrayDistance":        "300A00D9",
	"SourceToWedgeTrayDistance":           "300A00DA",
	"WedgeThinEdgePosition":               "300A00DB",
	"BolusID":                             "300A00DC",
	"BolusDescription":                    "300A00DD",
	"NumberOfCompensators":                "300A00E0",
	"MaterialID":                          "300A00E1",
	"TotalCompensatorTrayFactor":          "300A00E2",
	"CompensatorSequence":                 "300A00E3",
	"CompensatorNumber":                   "300A00E4",
	"CompensatorID":                       "300A00E5",
	"SourceToCompensatorTrayDistance":     "300A00E6",
	"CompensatorRows":                     "300A00E7",
	"CompensatorColumns":                  "300A00E8",
	"CompensatorPixelSpacing":             "300A00E9",
	"CompensatorPosition":                 "300A00EA",
	"CompensatorTransmissionData":         "300A00EB",
	"CompensatorThicknessData":            "300A00EC",
	"NumberOfBoli":                        "300A00ED",
	"CompensatorType":                     "300A00EE",
	"NumberOfBlocks":                      "300A00F0",
	"TotalBlockTrayFactor":                "300A00F2",
	"TotalBlockTrayWaterEquivThickness":   "300A00F3",
	"BlockSequence":                       "300A00F4",
	"BlockTrayID":                         "300A00F5",
	"SourceToBlockTrayDistance":           "300A00F6",
	"IsocenterToBlockTrayDistance":        "300A00F7",
	"BlockType":                           "300A00F8",
	"AccessoryCode":                       "300A00F9",
	"BlockDivergence":                     "300A00FA",
	"BlockMountingPosition":               "300A00FB",
	"BlockNumber":                         "300A00FC",
	"BlockName":                           "300A00FE",
	"BlockThickness":                      "300A0100",
	"BlockTransmission":                   "300A0102",
	"BlockNumberOfPoints":                 "300A0104",
	"BlockData":                           "300A0106",
	"ApplicatorSequence":                  "300A0107",
	"ApplicatorID":                        "300A0108",
	"ApplicatorType":                      "300A0109",
	"ApplicatorDescription":               "300A010A",
	"CumulativeDoseReferenceCoefficient":  "300A010C",
	"FinalCumulativeMetersetWeight":       "300A010E",
	"NumberOfControlPoints":               "300A0110",
	"ControlPointSequence":                "300A0111",
	"ControlPointIndex":                   "300A0112",
	"NominalBeamEnergy":                   "300A0114",
	"DoseRateSet":                         "300A0115",
	"WedgePositionSequence":               "300A0116",
	"WedgePosition":                       "300A0118",
	"BeamLimitingDevicePositionSequence":  "300A011A",
	"LeafJawPositions":                    "300A011C",
	"GantryAngle":                         "300A011E",
	"GantryRotationDirection":             "300A011F",
	"BeamLimitingDeviceAngle":             "300A0120",
	"BeamLimitingDeviceRotateDirection":   "300A0121",
	"PatientSupportAngle":                 "300A0122",
	"PatientSupportRotationDirection":     "300A0123",
	"TableTopEccentricAxisDistance":       "300A0124",
	"TableTopEccentricAngle":              "300A0125",
	"TableTopEccentricRotateDirection":    "300A0126",
	"TableTopVerticalPosition":            "300A0128",
	"TableTopLongitudinalPosition":        "300A0129",
	"TableTopLateralPosition":             "300A012A",
	"IsocenterPosition":                   "300A012C",
	"SurfaceEntryPoint":                   "300A012E",
	"SourceToSurfaceDistance":             "300A0130",
	"CumulativeMetersetWeight":            "300A0134",
	"TableTopPitchAngle":                  "300A0140",
	"TableTopPitchRotationDirection":      "300A0142",
	"TableTopRollAngle":                   "300A0144",
	"TableTopRollRotationDirection":       "300A0146",
	"HeadFixationAngle":                   "300A0148",
	"GantryPitchAngle":                    "300A014A",
	"GantryPitchRotationDirection":        "300A014C",
	"GantryPitchAngleTolerance":           "300A014E",
	"PatientSetupSequence":                "300A0180",
	"PatientSetupNumber":                  "300A0182",
	"PatientSetupLabel":                   "300A0183",
	"PatientAdditionalPosition":           "300A0184",
	"FixationDeviceSequence":              "300A0190",
	"FixationDeviceType":                  "300A0192",
	"FixationDeviceLabel":                 "300A0194",
	"FixationDeviceDescription":           "300A0196",
	"FixationDevicePosition":              "300A0198",
	"FixationDevicePitchAngle":            "300A0199",
	"FixationDeviceRollAngle":             "300A019A",
	"ShieldingDeviceSequence":             "300A01A0",
	"ShieldingDeviceType":                 "300A01A2",
	"ShieldingDeviceLabel":                "300A01A4",
	"ShieldingDeviceDescription":          "300A01A6",
	"ShieldingDevicePosition":             "300A01A8",
	"SetupTechnique":                      "300A01B0",
	"SetupTechniqueDescription":           "300A01B2",
	"SetupDeviceSequence":                 "300A01B4",
	"SetupDeviceType":                     "300A01B6",
	"SetupDeviceLabel":                    "300A01B8",
	"SetupDeviceDescription":              "300A01BA",
	"SetupDeviceParameter":                "300A01BC",
	"SetupReferenceDescription":           "300A01D0",
	"TableTopVerticalSetupDisplacement":   "300A01D2",
	"TableTopLongitudinalSetupDisplace":   "300A01D4",
	"TableTopLateralSetupDisplacement":    "300A01D6",
	"BrachyTreatmentTechnique":            "300A0200",
	"BrachyTreatmentType":                 "300A0202",
	"TreatmentMachineSequence":            "300A0206",
	"SourceSequence":                      "300A0210",
	"SourceNumber":                        "300A0212",
	"SourceType":                          "300A0214",
	"SourceManufacturer":                  "300A0216",
	"ActiveSourceDiameter":                "300A0218",
	"ActiveSourceLength":                  "300A021A",
	"SourceEncapsulationNomThickness":     "300A0222",
	"SourceEncapsulationNomTransmission":  "300A0224",
	"SourceIsotopeName":                   "300A0226",
	"SourceIsotopeHalfLife":               "300A0228",
	"SourceStrengthUnits":                 "300A0229",
	"ReferenceAirKermaRate":               "300A022A",
	"SourceStrength":                      "300A022B",
	"SourceStrengthReferenceDate":         "300A022C",
	"SourceStrengthReferenceTime":         "300A022E",
	"ApplicationSetupSequence":            "300A0230",
	"ApplicationSetupType":                "300A0232",
	"ApplicationSetupNumber":              "300A0234",
	"ApplicationSetupName":                "300A0236",
	"ApplicationSetupManufacturer":        "300A0238",
	"TemplateNumber":                      "300A0240",
	"TemplateType":                        "300A0242",
	"TemplateName":                        "300A0244",
	"TotalReferenceAirKerma":              "300A0250",
	"BrachyAccessoryDeviceSequence":       "300A0260",
	"BrachyAccessoryDeviceNumber":         "300A0262",
	"BrachyAccessoryDeviceID":             "300A0263",
	"BrachyAccessoryDeviceType":           "300A0264",
	"BrachyAccessoryDeviceName":           "300A0266",
	"BrachyAccessoryDeviceNomThickness":   "300A026A",
	"BrachyAccessoryDevNomTransmission":   "300A026C",
	"ChannelSequence":                     "300A0280",
	"ChannelNumber":                       "300A0282",
	"ChannelLength":                       "300A0284",
	"ChannelTotalTime":                    "300A0286",
	"SourceMovementType":                  "300A0288",
	"NumberOfPulses":                      "300A028A",
	"PulseRepetitionInterval":             "300A028C",
	"SourceApplicatorNumber":              "300A0290",
	"SourceApplicatorID":                  "300A0291",
	"SourceApplicatorType":                "300A0292",
	"SourceApplicatorName":                "300A0294",
	"SourceApplicatorLength":              "300A0296",
	"SourceApplicatorManufacturer":        "300A0298",
	"SourceApplicatorWallNomThickness":    "300A029C",
	"SourceApplicatorWallNomTrans":        "300A029E",
	"SourceApplicatorStepSize":            "300A02A0",
	"TransferTubeNumber":                  "300A02A2",
	"TransferTubeLength":                  "300A02A4",
	"ChannelShieldSequence":               "300A02B0",
	"ChannelShieldNumber":                 "300A02B2",
	"ChannelShieldID":                     "300A02B3",
	"ChannelShieldName":                   "300A02B4",
	"ChannelShieldNominalThickness":       "300A02B8",
	"ChannelShieldNominalTransmission":    "300A02BA",
	"FinalCumulativeTimeWeight":           "300A02C8",
	"BrachyControlPointSequence":          "300A02D0",
	"ControlPointRelativePosition":        "300A02D2",
	"ControlPoint3DPosition":              "300A02D4",
	"CumulativeTimeWeight":                "300A02D6",
	"CompensatorDivergence":               "300A02E0",
	"CompensatorMountingPosition":         "300A02E1",
	"SourceToCompensatorDistance":         "300A02E2",
	"TotalCompTrayWaterEquivThickness":    "300A02E3",
	"IsocenterToCompensatorTrayDistance":  "300A02E4",
	"CompensatorColumnOffset":             "300A02E5",
	"IsocenterToCompensatorDistances":     "300A02E6",
	"CompensatorRelStoppingPowerRatio":    "300A02E7",
	"CompensatorMillingToolDiameter":      "300A02E8",
	"IonRangeCompensatorSequence":         "300A02EA",
	"CompensatorDescription":              "300A02EB",
	"RadiationMassNumber":                 "300A0302",
	"RadiationAtomicNumber":               "300A0304",
	"RadiationChargeState":                "300A0306",
	"ScanMode":                            "300A0308",
	"VirtualSourceAxisDistances":          "300A030A",
	"SnoutSequence":                       "300A030C",
	"SnoutPosition":                       "300A030D",
	"SnoutID":                             "300A030F",
	"NumberOfRangeShifters":               "300A0312",
	"RangeShifterSequence":                "300A0314",
	"RangeShifterNumber":                  "300A0316",
	"RangeShifterID":                      "300A0318",
	"RangeShifterType":                    "300A0320",
	"RangeShifterDescription":             "300A0322",
	"NumberOfLateralSpreadingDevices":     "300A0330",
	"LateralSpreadingDeviceSequence":      "300A0332",
	"LateralSpreadingDeviceNumber":        "300A0334",
	"LateralSpreadingDeviceID":            "300A0336",
	"LateralSpreadingDeviceType":          "300A0338",
	"LateralSpreadingDeviceDescription":   "300A033A",
	"LateralSpreadingDevWaterEquivThick":  "300A033C",
	"NumberOfRangeModulators":             "300A0340",
	"RangeModulatorSequence":              "300A0342",
	"RangeModulatorNumber":                "300A0344",
	"RangeModulatorID":                    "300A0346",
	"RangeModulatorType":                  "300A0348",
	"RangeModulatorDescription":           "300A034A",
	"BeamCurrentModulationID":             "300A034C",
	"PatientSupportType":                  "300A0350",
	"PatientSupportID":                    "300A0352",
	"PatientSupportAccessoryCode":         "300A0354",
	"FixationLightAzimuthalAngle":         "300A0356",
	"FixationLightPolarAngle":             "300A0358",
	"MetersetRate":                        "300A035A",
	"RangeShifterSettingsSequence":        "300A0360",
	"RangeShifterSetting":                 "300A0362",
	"IsocenterToRangeShifterDistance":     "300A0364",
	"RangeShifterWaterEquivThickness":     "300A0366",
	"LateralSpreadingDeviceSettingsSeq":   "300A0370",
	"LateralSpreadingDeviceSetting":       "300A0372",
	"IsocenterToLateralSpreadingDevDist":  "300A0374",
	"RangeModulatorSettingsSequence":      "300A0380",
	"RangeModulatorGatingStartValue":      "300A0382",
	"RangeModulatorGatingStopValue":       "300A0384",
	"IsocenterToRangeModulatorDistance":   "300A038A",
	"ScanSpotTuneID":                      "300A0390",
	"NumberOfScanSpotPositions":           "300A0392",
	"ScanSpotPositionMap":                 "300A0394",
	"ScanSpotMetersetWeights":             "300A0396",
	"ScanningSpotSize":                    "300A0398",
	"NumberOfPaintings":                   "300A039A",
	"IonToleranceTableSequence":           "300A03A0",
	"IonBeamSequence":                     "300A03A2",
	"IonBeamLimitingDeviceSequence":       "300A03A4",
	"IonBlockSequence":                    "300A03A6",
	"IonControlPointSequence":             "300A03A8",
	"IonWedgeSequence":                    "300A03AA",
	"IonWedgePositionSequence":            "300A03AC",
	"ReferencedSetupImageSequence":        "300A0401",
	"SetupImageComment":                   "300A0402",
	"MotionSynchronizationSequence":       "300A0410",
	"ControlPointOrientation":             "300A0412",
	"GeneralAccessorySequence":            "300A0420",
	"GeneralAccessoryID":                  "300A0421",
	"GeneralAccessoryDescription":         "300A0422",
	"GeneralAccessoryType":                "300A0423",
	"GeneralAccessoryNumber":              "300A0424",
	"ReferencedRTPlanSequence":            "300C0002",
	"ReferencedBeamSequence":              "300C0004",
	"ReferencedBeamNumber":                "300C0006",
	"ReferencedReferenceImageNumber":      "300C0007",
	"StartCumulativeMetersetWeight":       "300C0008",
	"EndCumulativeMetersetWeight":         "300C0009",
	"ReferencedBrachyAppSetupSeq":         "300C000A",
	"ReferencedBrachyAppSetupNumber":      "300C000C",
	"ReferencedSourceNumber":              "300C000E",
	"ReferencedFractionGroupSequence":     "300C0020",
	"ReferencedFractionGroupNumber":       "300C0022",
	"ReferencedVerificationImageSeq":      "300C0040",
	"ReferencedReferenceImageSequence":    "300C0042",
	"ReferencedDoseReferenceSequence":     "300C0050",
	"ReferencedDoseReferenceNumber":       "300C0051",
	"BrachyReferencedDoseReferenceSeq":    "300C0055",
	"ReferencedStructureSetSequence":      "300C0060",
	"ReferencedPatientSetupNumber":        "300C006A",
	"ReferencedDoseSequence":              "300C0080",
	"ReferencedToleranceTableNumber":      "300C00A0",
	"ReferencedBolusSequence":             "300C00B0",
	"ReferencedWedgeNumber":               "300C00C0",
	"ReferencedCompensatorNumber":         "300C00D0",
	"ReferencedBlockNumber":               "300C00E0",
	"ReferencedControlPointIndex":         "300C00F0",
	"ReferencedControlPointSequence":      "300C00F2",
	"ReferencedStartControlPointIndex":    "300C00F4",
	"ReferencedStopControlPointIndex":     "300C00F6",
	"ReferencedRangeShifterNumber":        "300C0100",
	"ReferencedLateralSpreadingDevNum":    "300C0102",
	"ReferencedRangeModulatorNumber":      "300C0104",
	"ApprovalStatus":                      "300E0002",
	"ReviewDate":                          "300E0004",
	"ReviewTime":                          "300E0005",
	"ReviewerName":                        "300E0008",
	"TextGroupLength":                     "40000000",
	"Arbitrary":                           "40000010",
	"TextComments":                        "40004000",
	"ResultsID":                           "40080040",
	"ResultsIDIssuer":                     "40080042",
	"ReferencedInterpretationSequence":    "40080050",
	"InterpretationRecordedDate":          "40080100",
	"InterpretationRecordedTime":          "40080101",
	"InterpretationRecorder":              "40080102",
	"ReferenceToRecordedSound":            "40080103",
	"InterpretationTranscriptionDate":     "40080108",
	"InterpretationTranscriptionTime":     "40080109",
	"InterpretationTranscriber":           "4008010A",
	"InterpretationText":                  "4008010B",
	"InterpretationAuthor":                "4008010C",
	"InterpretationApproverSequence":      "40080111",
	"InterpretationApprovalDate":          "40080112",
	"InterpretationApprovalTime":          "40080113",
	"PhysicianApprovingInterpretation":    "40080114",
	"InterpretationDiagnosisDescription":  "40080115",
	"InterpretationDiagnosisCodeSeq":      "40080117",
	"ResultsDistributionListSequence":     "40080118",
	"DistributionName":                    "40080119",
	"DistributionAddress":                 "4008011A",
	"InterpretationID":                    "40080200",
	"InterpretationIDIssuer":              "40080202",
	"InterpretationTypeID":                "40080210",
	"InterpretationStatusID":              "40080212",
	"Impressions":                         "40080300",
	"ResultsComments":                     "40084000",
	"MACParametersSequence":               "4FFE0001",
	"CurveDimensions":                     "50xx0005",
	"NumberOfPoints":                      "50xx0010",
	"TypeOfData":                          "50xx0020",
	"CurveDescription":                    "50xx0022",
	"AxisUnits":                           "50xx0030",
	"AxisLabels":                          "50xx0040",
	"DataValueRepresentation":             "50xx0103",
	"MinimumCoordinateValue":              "50xx0104",
	"MaximumCoordinateValue":              "50xx0105",
	"CurveRange":                          "50xx0106",
	"CurveDataDescriptor":                 "50xx0110",
	"CoordinateStartValue":                "50xx0112",
	"CoordinateStepValue":                 "50xx0114",
	"CurveActivationLayer":                "50xx1001",
	"AudioType":                           "50xx2000",
	"AudioSampleFormat":                   "50xx2002",
	"NumberOfSamples":                     "50xx2006",
	"SampleRate":                          "50xx2008",
	"TotalTime":                           "50xx200A",
	"AudioSampleData":                     "50xx200C",
	"AudioComments":                       "50xx200E",
	"CurveLabel":                          "50xx2500",
	"ReferencedOverlayGroup":              "50xx2610",
	"CurveData":                           "50xx3000",
	"SharedFunctionalGroupsSequence":      "52009229",
	"PerFrameFunctionalGroupsSequence":    "52009230",
	"WaveformSequence":                    "54000100",
	"ChannelMinimumValue":                 "54000110",
	"ChannelMaximumValue":                 "54000112",
	"WaveformBitsAllocated":               "54001004",
	"WaveformSampleInterpretation":        "54001006",
	"WaveformPaddingValue":                "5400100A",
	"WaveformData":                        "54001010",
	"FirstOrderPhaseCorrectionAngle":      "56000010",
	"SpectroscopyData":                    "56000020",
	"OverlayGroupLength":                  "60000000",
	"OverlayRows":                         "60xx0010",
	"OverlayColumns":                      "60xx0011",
	"OverlayPlanes":                       "60xx0012",
	"NumberOfFramesInOverlay":             "60xx0015",
	"OverlayDescription":                  "60xx0022",
	"OverlayType":                         "60xx0040",
	"OverlaySubtype":                      "60xx0045",
	"OverlayOrigin":                       "60xx0050",
	"ImageFrameOrigin":                    "60xx0051",
	"OverlayPlaneOrigin":                  "60xx0052",
	"OverlayCompressionCode":              "60xx0060",
	"OverlayCompressionOriginator":        "60xx0061",
	"OverlayCompressionLabel":             "60xx0062",
	"OverlayCompressionDescription":       "60xx0063",
	"OverlayCompressionStepPointers":      "60xx0066",
	"OverlayRepeatInterval":               "60xx0068",
	"OverlayBitsGrouped":                  "60xx0069",
	"OverlayBitsAllocated":                "60xx0100",
	"OverlayBitPosition":                  "60xx0102",
	"OverlayFormat":                       "60xx0110",
	"OverlayLocation":                     "60xx0200",
	"OverlayCodeLabel":                    "60xx0800",
	"OverlayNumberOfTables":               "60xx0802",
	"OverlayCodeTableLocation":            "60xx0803",
	"OverlayBitsForCodeWord":              "60xx0804",
	"OverlayActivationLayer":              "60xx1001",
	"OverlayDescriptorGray":               "60xx1100",
	"OverlayDescriptorRed":                "60xx1101",
	"OverlayDescriptorGreen":              "60xx1102",
	"OverlayDescriptorBlue":               "60xx1103",
	"OverlaysGray":                        "60xx1200",
	"OverlaysRed":                         "60xx1201",
	"OverlaysGreen":                       "60xx1202",
	"OverlaysBlue":                        "60xx1203",
	"ROIArea":                             "60xx1301",
	"ROIMean":                             "60xx1302",
	"ROIStandardDeviation":                "60xx1303",
	"OverlayLabel":                        "60xx1500",
	"OverlayData":                         "60xx3000",
	"OverlayComments":                     "60xx4000",
	"PixelDataGroupLength":                "7Fxx0000",
	"PixelData":                           "7Fxx0010",
	"VariableNextDataGroup":               "7Fxx0011",
	"VariableCoefficientsSDVN":            "7Fxx0020",
	"VariableCoefficientsSDHN":            "7Fxx0030",
	"VariableCoefficientsSDDN":            "7Fxx0040",
	"DigitalSignaturesSequence":           "FFFAFFFA",
	"DataSetTrailingPadding":              "FFFCFFFC",
	"StartOfItem":                         "FFFEE000",
	"EndOfItems":                          "FFFEE00D",
	"EndOfSequence":                       "FFFEE0DD",
}
