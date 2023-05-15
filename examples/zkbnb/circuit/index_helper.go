/*
 * Copyright © 2022 ZkBNB Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package circuit

func AccountIndexToMerkleHelper(api API, accountIndex Variable) (merkleHelpers []Variable) {
	merkleHelpers = api.ToBinary(accountIndex, AccountMerkleLevels)
	return merkleHelpers
}

func AssetIdToMerkleHelper(api API, assetId Variable) (merkleHelpers []Variable) {
	merkleHelpers = api.ToBinary(assetId, AssetMerkleLevels)
	return merkleHelpers
}

func NftIndexToMerkleHelper(api API, nftIndex Variable) (merkleHelpers []Variable) {
	merkleHelpers = api.ToBinary(nftIndex, NftMerkleLevels)
	return merkleHelpers
}
