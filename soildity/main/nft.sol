// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
https://docs.openzeppelin.com/contracts/4.x/wizard

BaseUrl: 代幣 baseinfo
Mintable: 可挖掘
    AutoIncrementIds: NFT序號自動增長
Burnable: 可燃燒

**/


import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol"; // 
import "@openzeppelin/contracts/utils/Counters.sol"; // 計數器

contract MyToken is ERC721, ERC721URIStorage, ERC721Burnable, Ownable {
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdCounter;

    constructor() ERC721("MyToken", "MTK") {}

    function _baseURI() internal pure override returns (string memory) {
        return "http://qnimate.com/wp-content/uploads/2014/03/images2.jpg";
    }

    // 安全挖礦 onlyOwner (挖給那個,uri)
    function safeMint(address to, string memory uri) public onlyOwner {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment(); // + 1
        _safeMint(to, tokenId); // 挖 設置序列號
        _setTokenURI(tokenId, uri); // 設置token uri
    }

    // The following functions are overrides required by Solidity.
    // 燃燒
    function _burn(uint256 tokenId) internal override(ERC721, ERC721URIStorage) {
        super._burn(tokenId);
    }

    // 更具token 序列號 獲取NFT地址
    function tokenURI(uint256 tokenId)
    public
    view
    override(ERC721, ERC721URIStorage)
    returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
}