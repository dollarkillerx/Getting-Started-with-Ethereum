npm init --yes

npm install --save-dev hardhat

npx hardhat  // 初始化

npx hardhat node // 啓動節點

npx hardhat compile // 編譯

npx hardhat run .\scripts\deploy.js --network localhost    // 合約部署到本地

npx hardhat run .\tset\Lock.js --network localhost    // 單元測試