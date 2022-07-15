# Doc: https://docs.metamask.io/guide/

``` 
    <script>

        const isMetaMaskInstalled = ()=>{
            const { ethereum } = window;
            if (typeof ethereum === 'undefined') {
                return false;
            }

            return Boolean(ethereum && ethereum.isMetaMask);
        }

        const getAccount = async () => {
            try {
                // 不需要權限即可訪問
                const accounts = await ethereum.request({
                    method: 'eth_requestAccounts'
                });

                accountSpan.innerHTML = accounts;
            }catch (e) {
                console.error(e);
            }
        }

        const getNetworkAndChainID = async () => {
            try {
                const chainId = await ethereum.request({
                    method: 'eth_chainId'
                });

                chainIdSpan.innerHTML = chainId;

                const network = await ethereum.request({
                    method: 'net_version'
                });

                networkSpan.innerHTML = network;
            }catch (e) {
                console.error(e);
            }
        }

        const webInit = ()=> {
            const checkMetaMaskClient = async()=> {
                if (!isMetaMaskInstalled()) {
                    alert("Please install MetaMask");
                }else {
                    getNetworkAndChainID();
                    getAccount();
                }
            }

            checkMetaMaskClient();
        }

        getAccountButton.onclick = async () => {
            try {
                const accounts = await ethereum.request({
                    method: 'eth_accounts'
                });

                getAccountButtonResult.innerHTML = accounts;
            }catch (e) {
                console.error(e);
            }
        }

        sendButton.onclick = async () => {
            try {
                const accounts = await ethereum.request({
                    method: 'eth_accounts'
                });

                // 廢棄
                // web3.eth.sendTransaction({
                //     form: accounts[0],
                //     to: toAddr.value,
                //     value: 1000000000000000, // wei
                //     gas: 21000,
                //     gesPrice: 20000000000
                // },(result)=> {
                //     console.log(result)
                // })

                // 轉賬
                const transactionParameters = {
                    from: accounts[0],
                    to: toAddr.value,
                    value: '0x38d7ea4c68000',
                    gasPrice: '0x4a817c800',
                    gas: '0x5208',
                };

               try {
                   // txHash is a hex string
                   // As with any RPC call, it may throw an error
                   const txHash = await ethereum.request({
                       method: 'eth_sendTransaction',
                       params: [transactionParameters],
                   });

                   console.log(txHash)
               }catch (e) {
                   alert("轉賬失敗")
               }
            }catch (e) {
                console.error(e);
            }
        }

        // 監聽切換賬戶
        ethereum.on('accountsChanged',function (accounts) {
            console.log('accountsChanged');
            window.location.reload();
        })

        // 監聽切換網絡
        ethereum.on('chainChanged',function (accounts) {
            console.log('chainChanged');
            window.location.reload();
        })

        getBalanceButton.onclick = async () => {
            const accounts = await ethereum.request({
                method: 'eth_accounts'
            });

             ethereum.request({
                method: 'eth_getBalance',
                 params: [accounts[0],'latest']
            }).then(result => {
                getBalanceButtonResult.innerHTML = `${result} => ${parseInt(result,16)} wei`;
            }).catch(err=>{
                console.log(err)
             });

        }

        window.addEventListener('DOMContentLoaded',webInit)
    </script>
```