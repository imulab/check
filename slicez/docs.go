// slicez contains check.Step implementation related to slices.
//
// Particular element types are grouped under a namespace. For instance.
//
//	// Use namespace to access slice check.Step with respect to a specific type.
//	slicez.OfString.HasLength(5)
//	slicez.OfString.All(stringz.IsNotEmpty)
//
// Currently, only string slice is supported.
package slicez
