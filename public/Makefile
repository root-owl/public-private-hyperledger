.PHONY: generate clean run help upgrade-cc


generate:
	@./scripts/generate-certs.sh

clean:
	@./scripts/clean.sh

run:
	@cd ./chaincode && govendor sync
	@cd ..
	@./scripts/bootstrap-network.sh

help:
	@./scripts/help.sh

upgrade-cc:
ifndef CC_VERSION
	$(error CC_VERSION is not set)
endif
	@docker exec -e "CORE_PEER_ADDRESS=peer0.public.secc.com:7051" cli peer chaincode install -n secc -v $(CC_VERSION) -p secc/public/chaincode/
	@docker exec -e "CORE_PEER_ADDRESS=peer1.public.secc.com:7051" cli peer chaincode install -n secc -v $(CC_VERSION) -p secc/public/chaincode/
	@sleep 5
	@docker exec -e "CORE_PEER_ADDRESS=peer0.public.secc.com:7051" cli peer chaincode upgrade -o orderer0.secc.com:7050 -C public.secc -n secc -v $(CC_VERSION) -c '{"Args":[]}' -P "AND ('SECCPeerOrgMSP.peer')"

