package annotations

import "fmt"

// AnnotationName is a convenience type
type AnnotationName string

// Annotations is a convenience type
type Annotations map[string]string

// AnnotationPair represents a single annotation map entry
type AnnotationPair struct {
	name AnnotationName
	val  string
}

const baseAnnotationName = "ververicaplatform.fintechstudios.com"

var (
	ID                 = NewAnnotationName("id")
	Namespace          = NewAnnotationName("namespace")
	ResourceVersion    = NewAnnotationName("resource-version")
	DeploymentTargetID = NewAnnotationName("deployment-target-id")
	DeploymentID       = NewAnnotationName("deployment-id")
	JobID              = NewAnnotationName("job-id")
)

// NewAnnotationName creates a new annotation name with a given suffix
// with the common base "ververicaplatform.fintechstudios.com/"
func NewAnnotationName(key string) AnnotationName {
	return AnnotationName(fmt.Sprintf("%s/%s", baseAnnotationName, key))
}

// NewNamespacedAnnotationName creates a new annotation name with a given suffix
// with the common base "NAMESPACE.ververicaplatform.fintechstudios.com"
func NewNamespacedAnnotationName(subdomain, key string) AnnotationName {
	return AnnotationName(fmt.Sprintf("%s.%s/%s", subdomain, baseAnnotationName, key))
}

// Pair creates a new AnnotationPair
func Pair(attr AnnotationName, val string) AnnotationPair {
	return AnnotationPair{attr, val}
}

// Has determines whether or not an annotation is set
// safe to call on nil annotations
func Has(annotations Annotations, attr AnnotationName) bool {
	return Get(annotations, attr) != ""
}

// Get retrieves a single annotation, safe to call on nil annotations
func Get(annotations Annotations, attr AnnotationName) string {
	if annotations == nil {
		return ""
	}

	return annotations[string(attr)]
}

// Remove tries to remove an annotation and reports if it was present
func Remove(annotations Annotations, attr AnnotationName) bool {
	if !Has(annotations, attr) {
		return false
	}

	delete(annotations, string(attr))

	return true
}

// Set adds all the pairs to the annotations map
// safe to call on nil annotations
func Set(annotations Annotations, attrs ...AnnotationPair) Annotations {
	if annotations == nil {
		return Create(attrs...)
	}

	for _, attr := range attrs {
		annotations[string(attr.name)] = attr.val
	}
	return annotations
}

// Create makes a new Annotations map from a list of pairs
func Create(attrs ...AnnotationPair) Annotations {
	return Set(make(Annotations, len(attrs)), attrs...)
}

// EnsureExist either creates an empty annotation set or returns the given one
func EnsureExist(annotations Annotations) Annotations {
	if annotations != nil {
		return annotations
	}
	return Create()
}

// Copy creates a copy of the passed annotation set
func Copy(annotations Annotations) Annotations {
	if annotations == nil {
		return nil
	}
	target := Create()
	for k, v := range annotations {
		target[k] = v
	}
	return target
}
