// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node VariableSetStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("VariableSetStmt")
	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.FormatBool(node.IsLocal))
	ctx.WriteString(strconv.Itoa(int(node.Kind)))

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}