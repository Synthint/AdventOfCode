.PHONY: run
run:
	@cd $(Y)/Day_$(D) && \
	VAR="`ls -R | grep "main" | head -1`" && \
	echo "Found: $$VAR";\
	if [ $$VAR = "main.py" ]; then \
		echo "Detected Python For 12/$(D)/$(Y)"; \
		echo "Running...\n"; \
		python main.py; \
	fi; \
	if [ $$VAR = "main.rs" ]; then \
		echo "Detected Rust For 12/$(D)/$(Y)"; \
		echo "Running...\n"; \
		cargo run; \
	fi; \
	if [ $$VAR = "main.test.ts" ]; then \
		echo "Detected TypeScript For 12/$(D)/$(Y)"; \
		echo "\nRunning...\n"; \
		npx ts-node main.ts; \
	fi; \
	echo "\nFinished Run Job! \n\n"
	
.PHONY: test
test:
	@cd $(Y)/Day_$(D) && \
	VAR="`ls -R | grep "main" | head -1`" && \
	echo "Found: $$VAR";\
	if [ $$VAR = "main.py" ]; then \
		echo "Detected Python For 12/$(D)/$(Y)"; \
		echo "Testing...\n"; \
		pytest; \
	fi; \
	if [ $$VAR = "main.rs" ]; then \
		echo "Detected Rust For 12/$(D)/$(Y)"; \
		echo "Testing...\n"; \
		cargo test; \
	fi; \
	if [ $$VAR = "main.test.ts" ]; then \
		echo "Detected TypeScript For 12/$(D)/$(Y)"; \
		echo "\nTesting...\n"; \
		npx jest; \
	fi; \
	echo "\nFinished Test Job! \n\n"


.PHONY: run-and-test
run-and-test: run test
	
.PHONY: prepare
prepare:
	pip install pytest && \
	rustup install nightly && \
	rustup default nightly

.PHONY: clean
clean: 
	find . -type d -name 'target' -prune -exec rm -rf {} \; && \
	find . -type d -name '__pycache__' -prune -exec rm -rf {} \; && \
	find . -type d -name '.pytest_cache' -prune -exec rm -rf {} \;


.PHONY: gen-template
gen-template:
	@ mkdir -p ./${Y}/day_${D} && \
	if [ $(LANG) = "python" ]; then \
		echo "Creating Python Project For 12/$(D)/$(Y)"; \
		cp -r ./templates/$(LANG)/. ./$(Y)/Day_$(D)/ ; \
	fi; \
	if [ $(LANG) = "typescript" ]; then \
		echo "Creating TypeScript Project For 12/$(D)/$(Y)"; \
		cp -r ./templates/$(LANG)/. ./$(Y)/Day_$(D)/ ; \
		cd $(Y)/Day_$(D) && \
		npm i jest @types/jest ts-jest typescript ts-node -D; \
		echo {\"compilerOptions\":{\"esModuleInterop\":true}} > tsconfig.json ;\
		cat ../../templates/typescript/package.json > package.json ;\
	fi; \
	if [ $(LANG) = "rust" ]; then \
		echo "Creating Rust Project For 12/$(D)/$(Y)"; \
		cp -r ./templates/$(LANG)/. ./$(Y)/Day_$(D)/ ; \
		echo "Creating Cargo.toml"; \
		echo "[package]" > ./$(Y)/Day_$(D)/Cargo.toml ; \
		echo "name = \"day_$(D)\"" >> ./$(Y)/Day_$(D)/Cargo.toml ; \
		echo "version = \"0.1.0\"" >> ./$(Y)/Day_$(D)/Cargo.toml ; \
		echo "edition = \"2021\"" >> ./$(Y)/Day_$(D)/Cargo.toml ; \
		echo "[dependencies]" >> ./$(Y)/Day_$(D)/Cargo.toml ; \
	fi; \
	