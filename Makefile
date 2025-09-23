check_vegeta:
	@if  ! which vegeta; then \
    		echo "vegeta not found. Installing..."; \
    		brew install vegeta; \
	else \
    		echo "vegeta is already installed ✅"; \
    fi

attack: check_vegeta
	echo "PUT http://127.0.0.1:8080/v1/code-blocks" | vegeta attack -duration=${duration} --rate=${rate} | vegeta report --type=${type}
