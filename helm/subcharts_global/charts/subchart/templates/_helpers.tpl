{{- define "mychart.name" }}
   labels:
       app: familyproject
       k8s-app-name: famk8s
       owned_by: platform
{{- end }}

{{- define "mychart.version" }}
app_name: {{ $.Chart.Name }}
app_version: "{{ $.Chart.Version }}"
{{- end }}

