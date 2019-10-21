// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

const (

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExistsException".
	//
	// The resource with the name requested already exists.
	ErrCodeAlreadyExistsException = "AlreadyExistsException"

	// ErrCodeChangeSetNotFoundException for service response error code
	// "ChangeSetNotFound".
	//
	// The specified change set name or ID doesn't exit. To view valid change sets
	// for a stack, use the ListChangeSets action.
	ErrCodeChangeSetNotFoundException = "ChangeSetNotFound"

	// ErrCodeCreatedButModifiedException for service response error code
	// "CreatedButModifiedException".
	//
	// The specified resource exists, but has been changed.
	ErrCodeCreatedButModifiedException = "CreatedButModifiedException"

	// ErrCodeInsufficientCapabilitiesException for service response error code
	// "InsufficientCapabilitiesException".
	//
	// The template contains resources with capabilities that weren't specified
	// in the Capabilities parameter.
	ErrCodeInsufficientCapabilitiesException = "InsufficientCapabilitiesException"

	// ErrCodeInvalidChangeSetStatusException for service response error code
	// "InvalidChangeSetStatus".
	//
	// The specified change set can't be used to update the stack. For example,
	// the change set status might be CREATE_IN_PROGRESS, or the stack status might
	// be UPDATE_IN_PROGRESS.
	ErrCodeInvalidChangeSetStatusException = "InvalidChangeSetStatus"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// The specified operation isn't valid.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The quota for the resource has already been reached.
	//
	// For information on stack set limitations, see Limitations of StackSets (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-limitations.html).
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNameAlreadyExistsException for service response error code
	// "NameAlreadyExistsException".
	//
	// The specified name is already in use.
	ErrCodeNameAlreadyExistsException = "NameAlreadyExistsException"

	// ErrCodeOperationIdAlreadyExistsException for service response error code
	// "OperationIdAlreadyExistsException".
	//
	// The specified operation ID already exists.
	ErrCodeOperationIdAlreadyExistsException = "OperationIdAlreadyExistsException"

	// ErrCodeOperationInProgressException for service response error code
	// "OperationInProgressException".
	//
	// Another operation is currently in progress for this stack set. Only one operation
	// can be performed for a stack set at a given time.
	ErrCodeOperationInProgressException = "OperationInProgressException"

	// ErrCodeOperationNotFoundException for service response error code
	// "OperationNotFoundException".
	//
	// The specified ID refers to an operation that doesn't exist.
	ErrCodeOperationNotFoundException = "OperationNotFoundException"

	// ErrCodeStackInstanceNotFoundException for service response error code
	// "StackInstanceNotFoundException".
	//
	// The specified stack instance doesn't exist.
	ErrCodeStackInstanceNotFoundException = "StackInstanceNotFoundException"

	// ErrCodeStackSetNotEmptyException for service response error code
	// "StackSetNotEmptyException".
	//
	// You can't yet delete this stack set, because it still contains one or more
	// stack instances. Delete all stack instances from the stack set before deleting
	// the stack set.
	ErrCodeStackSetNotEmptyException = "StackSetNotEmptyException"

	// ErrCodeStackSetNotFoundException for service response error code
	// "StackSetNotFoundException".
	//
	// The specified stack set doesn't exist.
	ErrCodeStackSetNotFoundException = "StackSetNotFoundException"

	// ErrCodeStaleRequestException for service response error code
	// "StaleRequestException".
	//
	// Another operation has been performed on this stack set since the specified
	// operation was performed.
	ErrCodeStaleRequestException = "StaleRequestException"

	// ErrCodeTokenAlreadyExistsException for service response error code
	// "TokenAlreadyExistsException".
	//
	// A client request token already exists.
	ErrCodeTokenAlreadyExistsException = "TokenAlreadyExistsException"
)
