present:
	present -notes -base ./theme -content $(CONTENT)

update-theme:
	cp -rfv /home/dario/go/pkg/mod/golang.org/x/tools@v0.2.0/cmd/present/templates/ ./theme/
	cp -rfv /home/dario/go/pkg/mod/golang.org/x/tools@v0.2.0/cmd/present/static/ ./theme/