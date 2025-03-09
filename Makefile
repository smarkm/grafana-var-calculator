release:
	rm -rf smarkm-varcalculator-datasource/*
	cp -r dist/* smarkm-varcalculator-datasource/
	zip -r smarkm-varcalculator-datasource.v0.0.1.zip smarkm-varcalculator-datasource
	openssl dgst -sha1 smarkm-varcalculator-datasource.v0.0.1.zip > smarkm-varcalculator-datasource.v0.0.1.sha1 
copy:
	mv smarkm-varcalculator-datasource.v* /home/smark/