package pub

import (
	"context"
	"net/http"
)

// Actor represents ActivityPub's actor concept. It conceptually has an inbox
// and outbox that receives either a POST or GET request, which triggers side
// effects in the federating application.
//
// An Actor within an application may federate server-to-server (Federation
// Protocol), client-to-server (Social API), or both. The Actor represents the
// server in either use case.
//
// An actor can be created by calling NewSocialActor (only the Social Protocol
// is supported), NewFederatingActor (only the Federating Protocol is
// supported), NewActor (both are supported), or NewCustomActor (neither are).
//
// Not all Actors have the same behaviors depending on the constructor used to
// create them. Refer to the constructor's documentation to determine the exact
// behavior of the Actor on an application.
//
// The behaviors documented here are common to all Actors returned by any
// constructor.
type Actor interface {
	// PostInbox returns true if the request was handled as an ActivityPub
	// POST to an actor's inbox. If false, the request was not an
	// ActivityPub request and may still be handled by the caller in
	// another way, such as serving a web page.
	//
	// If the error is nil, then the ResponseWriter's headers and response
	// has already been written. If a non-nil error is returned, then no
	// response has been written.
	//
	// If the Actor was constructed with the Federated Protocol enabled,
	// side effects will occur.
	//
	// If the Federated Protocol is not enabled, writes the
	// http.StatusMethodNotAllowed status code in the response. No side
	// effects occur.
	PostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error)
	// GetInbox returns true if the request was handled as an ActivityPub
	// GET to an actor's inbox. If false, the request was not an ActivityPub
	// request and may still be handled by the caller in another way, such
	// as serving a web page.
	//
	// If the error is nil, then the ResponseWriter's headers and response
	// has already been written. If a non-nil error is returned, then no
	// response has been written.
	//
	// If the request is an ActivityPub request, the Actor will defer to the
	// application to determine the correct authorization of the request and
	// the resulting OrderedCollection to respond with. The Actor handles
	// serializing this OrderedCollection and responding with the correct
	// headers and http.StatusOK.
	GetInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error)
	// PostOutbox returns true if the request was handled as an ActivityPub
	// POST to an actor's outbox. If false, the request was not an
	// ActivityPub request and may still be handled by the caller in another
	// way, such as serving a web page.
	//
	// If the error is nil, then the ResponseWriter's headers and response
	// has already been written. If a non-nil error is returned, then no
	// response has been written.
	//
	// If the Actor was constructed with the Social Protocol enabled, side
	// effects will occur.
	//
	// If the Social Protocol is not enabled, writes the
	// http.StatusMethodNotAllowed status code in the response. No side
	// effects occur.
	//
	// If the Social and Federated Protocol are both enabled, it will handle
	// the side effects of receiving an ActivityStream Activity, and then
	// federate the Activity to peers.
	PostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error)
	// GetOutbox returns true if the request was handled as an ActivityPub
	// GET to an actor's outbox. If false, the request was not an
	// ActivityPub request.
	//
	// If the error is nil, then the ResponseWriter's headers and response
	// has already been written. If a non-nil error is returned, then no
	// response has been written.
	//
	// If the request is an ActivityPub request, the Actor will defer to the
	// application to determine the correct authorization of the request and
	// the resulting OrderedCollection to respond with. The Actor handles
	// serializing this OrderedCollection and responding with the correct
	// headers and http.StatusOK.
	GetOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error)
}
