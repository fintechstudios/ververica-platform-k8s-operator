package annotations

type AnnotationName string

type AnnotationPair struct {
	name AnnotationName
	val  string
}

const baseAnnotation = "ververicaplatform.fintechstudios.com/"

const (
	ID                 = AnnotationName(baseAnnotation + "id")
	Namespace          = AnnotationName(baseAnnotation + "namespace")
	ResourceVersion    = AnnotationName(baseAnnotation + "resource-version")
	DeploymentTargetID = AnnotationName(baseAnnotation + "deployment-target-id")
	DeploymentID       = AnnotationName(baseAnnotation + "deployment-id")
	JobID              = AnnotationName(baseAnnotation + "job-id")
)

func Pair(attr AnnotationName, val string) AnnotationPair {
	return AnnotationPair{attr, val}
}

func Has(annotations map[string]string, attr AnnotationName) bool {
	return Get(annotations, attr) != ""
}

func Get(annotations map[string]string, attr AnnotationName) string {
	if annotations == nil {
		return ""
	}

	return annotations[string(attr)]
}

func Set(annotations map[string]string, attrs ...AnnotationPair) map[string]string {
	for _, attr := range attrs {
		annotations[string(attr.name)] = attr.val
	}
	return annotations
}
