pragma solidity >0.8.0 < 0.9.0;

// test view and pure
contract TestContract {

	uint256 value;

	function store(uint256 number) public{
		value = number;
	}

	function retrieve() public view returns (uint256){
		return value;
	}
//
//	function greeting() public pure returns (string memory) {
//		return "Hello World";
//	}
}