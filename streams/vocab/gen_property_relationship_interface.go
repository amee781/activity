package vocab

import "net/url"

// RelationshipPropertyIteratorInterface represents a single value for the
// "relationship" property.
type RelationshipPropertyIteratorInterface interface {
	// Get returns the value of this property. When IsObject returns false,
	// Get will return any arbitrary value.
	Get() ObjectInterface
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// IsObject returns true if this property is set and not an IRI.
	IsObject() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o RelationshipPropertyIteratorInterface) bool
	// Name returns the name of this property: "relationship".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() RelationshipPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() RelationshipPropertyIteratorInterface
	// Set sets the value of this property. Calling IsObject afterwards will
	// return true.
	Set(v ObjectInterface)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}

// On a Relationship object, the relationship property identifies the kind of
// relationship that exists between subject and object.
//
// Example 140 (https://www.w3.org/TR/activitystreams-vocabulary/#ex22c-jsonld):
//   {
//     "object": {
//       "name": "John",
//       "type": "Person"
//     },
//     "relationship": "http://purl.org/vocab/relationship/acquaintanceOf",
//     "subject": {
//       "name": "Sally",
//       "type": "Person"
//     },
//     "summary": "Sally is an acquaintance of John's",
//     "type": "Relationship"
//   }
type RelationshipPropertyInterface interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "relationship"
	AppendIRI(v *url.URL)
	// AppendObject appends a Object value to the back of a list of the
	// property "relationship". Invalidates iterators that are traversing
	// using Prev.
	AppendObject(v ObjectInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) RelationshipPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() RelationshipPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() RelationshipPropertyIteratorInterface
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "relationship"
	// property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o RelationshipPropertyInterface) bool
	// Name returns the name of this property: "relationship".
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "relationship".
	PrependIRI(v *url.URL)
	// PrependObject prepends a Object value to the front of a list of the
	// property "relationship". Invalidates all iterators.
	PrependObject(v ObjectInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "relationship", regardless of its type. Panics if the
	// index is out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets a Object value to be at the specified index for the property
	// "relationship". Panics if the index is out of bounds. Invalidates
	// all iterators.
	Set(idx int, v ObjectInterface)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "relationship". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// Swap swaps the location of values at two indices for the "relationship"
	// property.
	Swap(i, j int)
}