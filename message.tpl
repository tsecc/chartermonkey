{{define "wazup"}}幹嘛~{{end}}
{{define "plusone"}}好喔, {{ .Name }} +1, 吱吱{{end}}
{{define "failed"}}那個...資料庫好像有點問題, 加不進去, 吱吱{{end}}
{{define "reject"}}建議不要私底下揪團~吱吱{{end}}
{{define "list"}}本週已+1的有{{range .}}{{.Namelist}}, {{end}}吱吱{{end}}
{{define "duplicate"}}誒...{{ .Name }}你已經在本週名單了哦~ 吱吱{{end}}