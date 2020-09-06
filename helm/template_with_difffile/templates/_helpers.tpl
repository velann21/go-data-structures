{{- define "mychart.name" }}
   labels:
       app: familyproject
       k8s-app-name: famk8s
       owned_by: platform
       servType: backend
       releaseName: "{{ $.Release.Name }}"
       chartName: "{{ $.Chart.Name }}"
       ReleaseSrv: "{{ $.Release.Service }}"
       ChartVersion: "{{ $.Chart.AppVersion }}"
{{- end }}

