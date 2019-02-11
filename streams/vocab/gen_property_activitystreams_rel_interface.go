package vocab

import "net/url"

// ActivityStreamsRelPropertyIterator represents a single value for the "rel"
// property.
type ActivityStreamsRelPropertyIterator interface {
	// Get returns the value of this property. When IsRFCRfc5988 returns
	// false, Get will return any arbitrary value.
	Get() string
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// IsRFCRfc5988 returns true if this property is set and not an IRI.
	IsRFCRfc5988() bool
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
	LessThan(o ActivityStreamsRelPropertyIterator) bool
	// Name returns the name of this property: "ActivityStreamsRel".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ActivityStreamsRelPropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ActivityStreamsRelPropertyIterator
	// Set sets the value of this property. Calling IsRFCRfc5988 afterwards
	// will return true.
	Set(v string)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}

// A link relation associated with a Link. The value MUST conform to both the
// [HTML5] and [RFC5988] "link relation" definitions. In the [HTML5], any
// string not containing the "space" U+0020, "tab" (U+0009), "LF" (U+000A),
// "FF" (U+000C), "CR" (U+000D) or "," (U+002C) characters can be used as a
// valid link relation.
//
// Example 131 (https://www.w3.org/TR/activitystreams-vocabulary/#ex149-jsonld):
//   {
//     "hreflang": "en",
//     "mediaType": "text/html",
//     "name": "Preview",
//     "rel": [
//       "canonical",
//       "preview"
//     ],
//     "type": "owl:Class",
//     "url": "http://example.org/abc"
//   }
type ActivityStreamsRelProperty interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "rel"
	AppendIRI(v *url.URL)
	// AppendRFCRfc5988 appends a rfc5988 value to the back of a list of the
	// property "rel". Invalidates iterators that are traversing using
	// Prev.
	AppendRFCRfc5988(v string)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ActivityStreamsRelPropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ActivityStreamsRelPropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ActivityStreamsRelPropertyIterator
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
	// Len returns the number of values that exist for the "rel" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsRelProperty) bool
	// Name returns the name of this property: "rel".
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "rel".
	PrependIRI(v *url.URL)
	// PrependRFCRfc5988 prepends a rfc5988 value to the front of a list of
	// the property "rel". Invalidates all iterators.
	PrependRFCRfc5988(v string)
	// Remove deletes an element at the specified index from a list of the
	// property "rel", regardless of its type. Panics if the index is out
	// of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets a rfc5988 value to be at the specified index for the property
	// "rel". Panics if the index is out of bounds. Invalidates all
	// iterators.
	Set(idx int, v string)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "rel". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// Swap swaps the location of values at two indices for the "rel" property.
	Swap(i, j int)
}