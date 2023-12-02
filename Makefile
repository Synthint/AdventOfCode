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
		npm i jest @types/jest ts-jest typescript -D; \
		npm install ts-node -D; \
		echo {\"compilerOptions\":{\"esModuleInterop\":true}} > tsconfig.json ;\
		echo "\nRunning...\n"; \
		npx ts-node main.ts; \
	fi; \
	echo "Finished Run Job! \n\n"
	
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
		npm i jest @types/jest ts-jest typescript ts-node -D; \
		echo {\"compilerOptions\":{\"esModuleInterop\":true}} > tsconfig.json ;\
		echo "\nTesting...\n"; \
		npx jest; \
	fi; \
	echo "Finished Test Job! \n\n"


.PHONY: run-and-test
run-and-test: run test
	
.PHONY: prepare
prepare:
	pip install pytest

.PHONY: clean
clean: 
	find . -type d -name 'target' -prune -exec rm -rf {} \; && \
	find . -type d -name '__pycache__' -prune -exec rm -rf {} \; && \
	find . -type d -name '.pytest_cache' -prune -exec rm -rf {} \;
