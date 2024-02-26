package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dymensionxyz/dymension/v3/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgSubmitFraud_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmitFraud
		err  error
	}{
		{
			name: "valid",
			msg: MsgSubmitFraud{
				Creator:    sample.AccAddress(),
				RollappID:  "rollapp_id",
				FraudProof: fpwithtx,
			},
		}, {
			name: "invalid address",
			msg: MsgSubmitFraud{
				Creator:    "invalid_address",
				RollappID:  "rollapp_id",
				FraudProof: fpwithtx,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "empty rollapp ID",
			msg: MsgSubmitFraud{
				Creator:    sample.AccAddress(),
				RollappID:  "",
				FraudProof: fpwithtx,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid fraud proof",
			msg: MsgSubmitFraud{
				Creator:    sample.AccAddress(),
				FraudProof: "invalid_fraud_proof",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

const (
	fpwithtx = `{
		"block_height": 85,
		"pre_state_app_hash": "/eXNbrvakhf2HKjsQ/OWwjT+5+vj/gqevNd5UchMeAc=",
		"expected_valid_app_hash": "Ahu/9ezjIucqYh10UsfzZmo3IQdYV4XIIyaU788FpcA=",
		"state_witness": {
		  "acc": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "YWNj",
			  "data": "Cv8BCgNhY2MSIAW4sPD1tOmXgbBwT/MNKC3mk/HCie9CvBY6g8GgOO1zGgkIARgBIAEqAQAiJwgBEgEBGiAJqB27FyhoLIGFO5kuvr1abT3muOcwPocX4h+HlMEyViInCAESAQEaIF1ED+gEy67hizY8itScXy9s7DlEJ+X9A4kHUB2jxh/KIicIARIBARog2fRYFBiPTQs0J5VN+27ohgNv71CZvNkjM/89anDRcn4iJwgBEgEBGiAedVgSOdzQZvfWW7RtE7jmfcB6g25IAq6vBrcqLIzjFiInCAESAQEaIJn5Nm4cA+0ZJ/lYo8IvsbqfWKWcrc5SCElDTL4Uv8sY"
			},
			"root_hash": "Bbiw8PW06ZeBsHBP8w0oLeaT8cKJ70K8FjqDwaA47XM=",
			"witness_data": [
			  {
				"operation": 1,
				"key": "ATrJQkUZo+lJ3HFy9P2CMhT6ccaV",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "YWNjb3VudE51bWJlcgAAAAAAAAAA",
					"data": "CvEBChVhY2NvdW50TnVtYmVyAAAAAAAAAAASFDrJQkUZo+lJ3HFy9P2CMhT6ccaVGgwIARgBIAEqBAACqgEiLAgBEgUCBKoBIBohIFhm2QKWHcimSTDAOOdGvhHatLwtf8zdR9WzBjronveVIiwIARIFBAiqASAaISA8vncyIITI+qA5JKLGiGo8fd6wk+/Yr0w9++OPYD/J3SIsCAESBQgYqgEgGiEgEnsezMFuL8j/CYbghIHW+6f4DY1E11kWLcFoSQq7pVUiKggBEiYKLqoBICVIPm94qBQv0ioyDOi02YfFL0bSHtVnTnnizh3eJ/ysIA=="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AXtf4itURvfGLqJ7i9cc75TgPz3y",
					"data": "Cr4CChUBe1/iK1RG98YuonuL1xzvlOA/PfISZAoiL2Nvc21vcy5hdXRoLnYxYmV0YTEuTW9kdWxlQWNjb3VudBI+Ci8KK2V0aG0xMGQwN3kyNjVnbW11dnQ0ejB3OWF3ODgwam5zcjcwMGpwdmE4NDMYBhIDZ292GgZidXJuZXIaCwgBGAEgASoDAAICIisIARIEBAYCIBohIJ5cTp0F8NIQW7Bi9OoyyaO9RBO4uRYvR7y2OUIw/qx9IisIARIEBgoCIBohIEgLrKUPDLxgN0Nz6q8SR7Mrx33Dy8dcuW2cUQwroPE0IioIARImCBaqASD7CpWxCtSbxRe++8CmuDoV+uJm+ZPZCZkhEC7OPcOGPSAiLAgBEgUKLqoBIBohINxVKXAP4KFIum02toFmrSGUxzMqjY+MHVdustNbDeBc"
				  },
				  {
					"type": "ics23:iavl",
					"key": "AVkRuETXvCJGVP4NzRa6vS0lPy/f",
					"data": "Ct0CChUBWRG4RNe8IkZU/g3NFrq9LSU/L98SgAEKIi9jb3Ntb3MuYXV0aC52MWJldGExLk1vZHVsZUFjY291bnQSWgovCitldGhtMXR5Z21zM3hoaHMzeXY0ODdwaHgzZHc0YTk1am43dDdsNjRtdXZwGAUSFm5vdF9ib25kZWRfdG9rZW5zX3Bvb2waBmJ1cm5lchoHc3Rha2luZxoLCAEYASABKgMAAgIiLAgBEgUEBqgBIBohIPTIcudGVx58bqw6ugYSMobIJ3oe4tpPp3QttdhRmH0uIioIARImBgyqASBN2hxYJ/pDXA4idfwOco6wSrFtsVFRFd5KcfxqajK5XyAiLAgBEgUIFqoBIBohILrjlbjprB2W7BXbPn8cV/h7mK2oNOio23Wnm56wZF5gIiwIARIFCi6qASAaISDcVSlwD+ChSLptNraBZq0hlMczKo2PjB1XbrLTWw3gXA=="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AU/qdkJ7g0WGHoCjVAqKnZNv05OR",
					"data": "CtgCChUBT+p2QnuDRYYegKNUCoqdk2/Tk5ESfAoiL2Nvc21vcy5hdXRoLnYxYmV0YTEuTW9kdWxlQWNjb3VudBJWCi8KK2V0aG0xZmw0OHZzbm1zZHpjdjg1cTVkMnE0ejVhamRoYTh5dTN3NDhkNjQYBBISYm9uZGVkX3Rva2Vuc19wb29sGgZidXJuZXIaB3N0YWtpbmcaCwgBGAEgASoDAAICIioIARImBAaqASBKLDKpFlfO3lUZ2MN6aC57eUKMtxUOQYMGGAzeDTI6nSAiLAgBEgUGDKoBIBohIH+Z67vw4LRnUJgKW4ux55U6evZINX43eHVGjwgPn305IiwIARIFCBaqASAaISC645W46awdluwV2z5/HFf4e5itqDToqNt1p5uesGReYCIsCAESBQouqgEgGiEg3FUpcA/goUi6bTa2gWatIZTHMyqNj4wdV26y01sN4Fw="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AUfusurDUOGSO4y9+kOWoHezbmKg",
					"data": "CvgCChUBR+6y6sNQ4ZI7jL36Q5agd7NuYqASbgoiL2Nvc21vcy5hdXRoLnYxYmV0YTEuTW9kdWxlQWNjb3VudBJICi8KK2V0aG0xZ2xodDk2a3IycnNleXd1dmhoYXk4OTRxdzdla3VjNHEzMmFjMnkYCBIFZXJjMjAaBm1pbnRlchoGYnVybmVyGgsIARgBIAEqAwACAiIqCAESJgIEqgEgHLyp48CosH93LXuQalgdVuQ8mSzGs4pP/Lh6l43DLIMgIiwIARIFBAaqASAaISDOQdKmW+OaEL7czhRtvSnPmc3OnqCDbEswX/rp/O7ClSIsCAESBQYMqgEgGiEgf5nru/DgtGdQmApbi7HnlTp69kg1fjd4dUaPCA+ffTkiLAgBEgUIFqoBIBohILrjlbjprB2W7BXbPn8cV/h7mK2oNOio23Wnm56wZF5gIiwIARIFCi6qASAaISDcVSlwD+ChSLptNraBZq0hlMczKo2PjB1XbrLTWw3gXA=="
				  },
				  {
					"type": "ics23:iavl",
					"key": "ATrJQkUZo+lJ3HFy9P2CMhT6ccaV",
					"data": "CvgDChUBOslCRRmj6UnccXL0/YIyFPpxxpUS6gEKHi9ldGhlcm1pbnQudHlwZXMudjEuRXRoQWNjb3VudBLHAQqAAQorZXRobTE4dHk1eTNnZTUwNTVuaHIzd3Q2MG1xM2p6bmE4cjM1NGU0ajhmYxJPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQJn8mJfKSWDNyaY7onyhXJN3ZVUVxs6TSxX7j4/gIuoQSADEkIweGM1ZDI0NjAxODZmNzIzM2M5MjdlN2RiMmRjYzcwM2MwZTUwMGI2NTNjYTgyMjczYjdiZmFkODA0NWQ4NWE0NzAaDAgBGAEgASoEAAKqASIsCAESBQIEqgEgGiEgH9CPkmmCFBO0IvPyxxEW/CD3tWMUjzjRC/QU/1MLV3ciLAgBEgUEBqoBIBohIM5B0qZb45oQvtzOFG29Kc+Zzc6eoINsSzBf+un87sKVIiwIARIFBgyqASAaISB/meu78OC0Z1CYCluLseeVOnr2SDV+N3h1Ro8ID599OSIsCAESBQgWqgEgGiEguuOVuOmsHZbsFds+fxxX+HuYrag06KjbdaebnrBkXmAiLAgBEgUKLqoBIBohINxVKXAP4KFIum02toFmrSGUxzMqjY+MHVdustNbDeBc"
				  }
				]
			  },
			  {
				"operation": 1,
				"key": "AXQgVx8nLNdp+jNVP71j3Px1xR1Y",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "YWNjb3VudE51bWJlcgAAAAAAAAAA",
					"data": "CvEBChVhY2NvdW50TnVtYmVyAAAAAAAAAAASFDrJQkUZo+lJ3HFy9P2CMhT6ccaVGgwIARgBIAEqBAACqgEiLAgBEgUCBKoBIBohIFhm2QKWHcimSTDAOOdGvhHatLwtf8zdR9WzBjronveVIiwIARIFBAiqASAaISA8vncyIITI+qA5JKLGiGo8fd6wk+/Yr0w9++OPYD/J3SIsCAESBQgYqgEgGiEgEnsezMFuL8j/CYbghIHW+6f4DY1E11kWLcFoSQq7pVUiKggBEiYKLqoBICVIPm94qBQv0ioyDOi02YfFL0bSHtVnTnnizh3eJ/ysIA=="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AXtf4itURvfGLqJ7i9cc75TgPz3y",
					"data": "Cr4CChUBe1/iK1RG98YuonuL1xzvlOA/PfISZAoiL2Nvc21vcy5hdXRoLnYxYmV0YTEuTW9kdWxlQWNjb3VudBI+Ci8KK2V0aG0xMGQwN3kyNjVnbW11dnQ0ejB3OWF3ODgwam5zcjcwMGpwdmE4NDMYBhIDZ292GgZidXJuZXIaCwgBGAEgASoDAAICIisIARIEBAYCIBohIJ5cTp0F8NIQW7Bi9OoyyaO9RBO4uRYvR7y2OUIw/qx9IisIARIEBgoCIBohIEgLrKUPDLxgN0Nz6q8SR7Mrx33Dy8dcuW2cUQwroPE0IioIARImCBaqASD7CpWxCtSbxRe++8CmuDoV+uJm+ZPZCZkhEC7OPcOGPSAiLAgBEgUKLqoBIBohINxVKXAP4KFIum02toFmrSGUxzMqjY+MHVdustNbDeBc"
				  },
				  {
					"type": "ics23:iavl",
					"key": "AVkRuETXvCJGVP4NzRa6vS0lPy/f",
					"data": "Ct0CChUBWRG4RNe8IkZU/g3NFrq9LSU/L98SgAEKIi9jb3Ntb3MuYXV0aC52MWJldGExLk1vZHVsZUFjY291bnQSWgovCitldGhtMXR5Z21zM3hoaHMzeXY0ODdwaHgzZHc0YTk1am43dDdsNjRtdXZwGAUSFm5vdF9ib25kZWRfdG9rZW5zX3Bvb2waBmJ1cm5lchoHc3Rha2luZxoLCAEYASABKgMAAgIiLAgBEgUEBqgBIBohIPTIcudGVx58bqw6ugYSMobIJ3oe4tpPp3QttdhRmH0uIioIARImBgyqASBN2hxYJ/pDXA4idfwOco6wSrFtsVFRFd5KcfxqajK5XyAiLAgBEgUIFqoBIBohILrjlbjprB2W7BXbPn8cV/h7mK2oNOio23Wnm56wZF5gIiwIARIFCi6qASAaISDcVSlwD+ChSLptNraBZq0hlMczKo2PjB1XbrLTWw3gXA=="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AW8hbgKzekTtRBgxnwg9q+NEAPkz",
					"data": "CqADChUBbyFuArN6RO1EGDGfCD2r40QA+TMSlwEKHi9ldGhlcm1pbnQudHlwZXMudjEuRXRoQWNjb3VudBJ1Ci8KK2V0aG0xZHVza3VxNG4wZnp3NjNxY3h4MHNzMGR0dWR6cXA3Zm5zMHMwN3gYARJCMHhjNWQyNDYwMTg2ZjcyMzNjOTI3ZTdkYjJkY2M3MDNjMGU1MDBiNjUzY2E4MjI3M2I3YmZhZDgwNDVkODVhNDcwGgsIARgBIAEqAwACAiIsCAESBQIEqAEgGiEg7kbhLy8uWZqLx/4fWpxx5/tNITMkjHhwmnesuSKlh/wiKggBEiYEBqgBIGhK702xTBv4gsnOtv92Hhnhl5Fo2lV3C9aNmKipJWOQICIqCAESJgYMqgEgTdocWCf6Q1wOInX8DnKOsEqxbbFRURXeSnH8amoyuV8gIiwIARIFCBaqASAaISC645W46awdluwV2z5/HFf4e5itqDToqNt1p5uesGReYCIsCAESBQouqgEgGiEg3FUpcA/goUi6bTa2gWatIZTHMyqNj4wdV26y01sN4Fw="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AXQgVx8nLNdp+jNVP71j3Px1xR1Y",
					"data": "Cp8DChUBdCBXHycs12n6M1U/vWPc/HXFHVgSlwEKHi9ldGhlcm1pbnQudHlwZXMudjEuRXRoQWNjb3VudBJ1Ci8KK2V0aG0xd3NzOXc4ZTg5bnRrbjczbjI1bG02Yzd1bDM2dTI4MmM0c3E1cW0YChJCMHhjNWQyNDYwMTg2ZjcyMzNjOTI3ZTdkYjJkY2M3MDNjMGU1MDBiNjUzY2E4MjI3M2I3YmZhZDgwNDVkODVhNDcwGgwIARgBIAEqBAACqAEiKggBEiYCBKgBIPjN9PCPZMatdeI1VTi5BhgwynzNGGNTaZBcCRcStR9sICIqCAESJgQGqAEgaErvTbFMG/iCyc62/3YeGeGXkWjaVXcL1o2YqKklY5AgIioIARImBgyqASBN2hxYJ/pDXA4idfwOco6wSrFtsVFRFd5KcfxqajK5XyAiLAgBEgUIFqoBIBohILrjlbjprB2W7BXbPn8cV/h7mK2oNOio23Wnm56wZF5gIiwIARIFCi6qASAaISDcVSlwD+ChSLptNraBZq0hlMczKo2PjB1XbrLTWw3gXA=="
				  }
				]
			  }
			]
		  },
		  "bank": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "YmFuaw==",
			  "data": "Cv4BCgRiYW5rEiCkk8h++gbQdQZ3PRlvOSavVR8lF2PrQs3AGoMk2FafohoJCAEYASABKgEAIiUIARIhAdz7dJLe1ht+mlFERhChTzzMER6cGGLNcNfiW9/hawwiIicIARIBARogXUQP6ATLruGLNjyK1JxfL2zsOUQn5f0DiQdQHaPGH8oiJwgBEgEBGiDZ9FgUGI9NCzQnlU37buiGA2/vUJm82SMz/z1qcNFyfiInCAESAQEaIB51WBI53NBm99ZbtG0TuOZ9wHqDbkgCrq8GtyosjOMWIicIARIBARogmfk2bhwD7Rkn+Vijwi+xup9YpZytzlIISUNMvhS/yxg="
			},
			"root_hash": "pJPIfvoG0HUGdz0Zbzkmr1UfJRdj60LNwBqDJNhWn6I=",
			"witness_data": [
			  {
				"operation": 1,
				"key": "AhQ6yUJFGaPpSdxxcvT9gjIU+nHGlWFkdW0=",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "AhRP6nZCe4NFhh6Ao1QKip2Tb9OTkWFkdW0=",
					"data": "CswBChoCFE/qdkJ7g0WGHoCjVAqKnZNv05ORYWR1bRIbNTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGgsIARgBIAEqAwACAiIqCAESJgIEqgEgBKs6C4rEQg4/tq3p4VCuTsYQy+WsB8oiswuUasLj1lEgIioIARImBAiqASBs0IEDDkQuoiP16tHyjeQyYCy7GQrh5AwHMr9u0VqOKyAiLAgBEgUIFKoBIBohIJS2u0QtVxEwt9cJS1yaH2d+4cnKUL1eijyfZSFmxajX"
				  },
				  {
					"type": "ics23:iavl",
					"key": "AhRvIW4Cs3pE7UQYMZ8IPavjRAD5M2FkdW0=",
					"data": "CvwBChoCFG8hbgKzekTtRBgxnwg9q+NEAPkzYWR1bRIbMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGgsIARgBIAEqAwACAiIsCAESBQIEqgEgGiEgTKCBvg0DQmIZjd1G9X0/z1PqesM6xLasc6hssDm9fFMiLAgBEgUEBqoBIBohIB++wMEwc7l/SYUEk+aRgmW9SNBg7TuLfMwGXVzo1SbzIiwIARIFBgyqASAaISBk0r+ufEOmlWtaxsEWckD1gPh2kHL0MdheTXmPAQRPwyIqCAESJggUqgEgRrFFqDf89LOgeN/QSxNIKJqLRx3Tw/FYDk7xQ1C7UKUg"
				  },
				  {
					"type": "ics23:iavl",
					"key": "AhQ6yUJFGaPpSdxxcvT9gjIU+nHGlWFkdW0=",
					"data": "Cs8BChoCFDrJQkUZo+lJ3HFy9P2CMhT6ccaVYWR1bRIbMzk5OTk5OTk5OTk5OTk5OTk5OTk5ODAwMDAwGgwIARgBIAEqBAACqgEiLAgBEgUCBKoBIBohIKU8yTwbJN1z9LAsRHsw6vPagW2WjSIdfU1GFPJvQA8wIioIARImBAiqASBs0IEDDkQuoiP16tHyjeQyYCy7GQrh5AwHMr9u0VqOKyAiLAgBEgUIFKoBIBohIJS2u0QtVxEwt9cJS1yaH2d+4cnKUL1eijyfZSFmxajX"
				  }
				]
			  },
			  {
				"operation": 1,
				"key": "A2FkdW0AFDrJQkUZo+lJ3HFy9P2CMhT6ccaV",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFE/qdkJ7g0WGHoCjVAqKnZNv05OR",
					"data": "CrMBChsDYWR1bQAUT+p2QnuDRYYegKNUCoqdk2/Tk5ESAQAaCwgBGAEgASoDAAICIiwIARIFBAaoASAaISBP3asyRmeyidQGCOiCsJrrkx+sMJGLG40hlUxhSs84ziIqCAESJgYMqgEgfRvh8Dp4twQb4atKiTLih8ERLofVSBKl4pOrITaPmrsgIioIARImCBSqASBGsUWoN/z0s6B439BLE0gomotHHdPD8VgOTvFDULtQpSA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFDrJQkUZo+lJ3HFy9P2CMhT6ccaV",
					"data": "CrMBChsDYWR1bQAUOslCRRmj6UnccXL0/YIyFPpxxpUSAQAaCwgBGAEgASoDAAICIioIARImBAaqASDIHeINewrAgcTbwOGzbDiNdABmdksZt97BfMrlrspPEiAiLAgBEgUGDKoBIBohIGTSv658Q6aVa1rGwRZyQPWA+HaQcvQx2F5NeY8BBE/DIioIARImCBSqASBGsUWoN/z0s6B439BLE0gomotHHdPD8VgOTvFDULtQpSA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AhRvIW4Cs3pE7UQYMZ8IPavjRAD5M2FkdW0=",
					"data": "CvwBChoCFG8hbgKzekTtRBgxnwg9q+NEAPkzYWR1bRIbMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGgsIARgBIAEqAwACAiIsCAESBQIEqgEgGiEgTKCBvg0DQmIZjd1G9X0/z1PqesM6xLasc6hssDm9fFMiLAgBEgUEBqoBIBohIB++wMEwc7l/SYUEk+aRgmW9SNBg7TuLfMwGXVzo1SbzIiwIARIFBgyqASAaISBk0r+ufEOmlWtaxsEWckD1gPh2kHL0MdheTXmPAQRPwyIqCAESJggUqgEgRrFFqDf89LOgeN/QSxNIKJqLRx3Tw/FYDk7xQ1C7UKUg"
				  }
				]
			  },
			  {
				"operation": 1,
				"key": "AhR0IFcfJyzXafozVT+9Y9z8dcUdWGFkdW0=",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "AhR0IFcfJyzXafozVT+9Y9z8dcUdWGFkdW0=",
					"data": "CuYBChoCFHQgVx8nLNdp+jNVP71j3Px1xR1YYWR1bRIGMjAwMDAwGgwIARgBIAEqBAACqgEiKggBEiYCBKoBIFrzkvkooK9Y1DelQr/gZmzgOa8D0tP2w0CnbIZWHJfvICIsCAESBQQGqgEgGiEgH77AwTBzuX9JhQST5pGCZb1I0GDtO4t8zAZdXOjVJvMiLAgBEgUGDKoBIBohIGTSv658Q6aVa1rGwRZyQPWA+HaQcvQx2F5NeY8BBE/DIioIARImCBSqASBGsUWoN/z0s6B439BLE0gomotHHdPD8VgOTvFDULtQpSA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AhRvIW4Cs3pE7UQYMZ8IPavjRAD5M2FkdW0=",
					"data": "CvwBChoCFG8hbgKzekTtRBgxnwg9q+NEAPkzYWR1bRIbMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGgsIARgBIAEqAwACAiIsCAESBQIEqgEgGiEgTKCBvg0DQmIZjd1G9X0/z1PqesM6xLasc6hssDm9fFMiLAgBEgUEBqoBIBohIB++wMEwc7l/SYUEk+aRgmW9SNBg7TuLfMwGXVzo1SbzIiwIARIFBgyqASAaISBk0r+ufEOmlWtaxsEWckD1gPh2kHL0MdheTXmPAQRPwyIqCAESJggUqgEgRrFFqDf89LOgeN/QSxNIKJqLRx3Tw/FYDk7xQ1C7UKUg"
				  },
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFE/qdkJ7g0WGHoCjVAqKnZNv05OR",
					"data": "CrMBChsDYWR1bQAUT+p2QnuDRYYegKNUCoqdk2/Tk5ESAQAaCwgBGAEgASoDAAICIiwIARIFBAaoASAaISBP3asyRmeyidQGCOiCsJrrkx+sMJGLG40hlUxhSs84ziIqCAESJgYMqgEgfRvh8Dp4twQb4atKiTLih8ERLofVSBKl4pOrITaPmrsgIioIARImCBSqASBGsUWoN/z0s6B439BLE0gomotHHdPD8VgOTvFDULtQpSA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFDrJQkUZo+lJ3HFy9P2CMhT6ccaV",
					"data": "CrMBChsDYWR1bQAUOslCRRmj6UnccXL0/YIyFPpxxpUSAQAaCwgBGAEgASoDAAICIioIARImBAaqASDIHeINewrAgcTbwOGzbDiNdABmdksZt97BfMrlrspPEiAiLAgBEgUGDKoBIBohIGTSv658Q6aVa1rGwRZyQPWA+HaQcvQx2F5NeY8BBE/DIioIARImCBSqASBGsUWoN/z0s6B439BLE0gomotHHdPD8VgOTvFDULtQpSA="
				  }
				]
			  },
			  {
				"operation": 1,
				"key": "A2FkdW0AFHQgVx8nLNdp+jNVP71j3Px1xR1Y",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "AhRvIW4Cs3pE7UQYMZ8IPavjRAD5M2FkdW0=",
					"data": "CvwBChoCFG8hbgKzekTtRBgxnwg9q+NEAPkzYWR1bRIbMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGgsIARgBIAEqAwACAiIsCAESBQIEqgEgGiEgTKCBvg0DQmIZjd1G9X0/z1PqesM6xLasc6hssDm9fFMiLAgBEgUEBqoBIBohIB++wMEwc7l/SYUEk+aRgmW9SNBg7TuLfMwGXVzo1SbzIiwIARIFBgyqASAaISBk0r+ufEOmlWtaxsEWckD1gPh2kHL0MdheTXmPAQRPwyIqCAESJggUqgEgRrFFqDf89LOgeN/QSxNIKJqLRx3Tw/FYDk7xQ1C7UKUg"
				  },
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFE/qdkJ7g0WGHoCjVAqKnZNv05OR",
					"data": "CrMBChsDYWR1bQAUT+p2QnuDRYYegKNUCoqdk2/Tk5ESAQAaCwgBGAEgASoDAAICIiwIARIFBAaoASAaISBP3asyRmeyidQGCOiCsJrrkx+sMJGLG40hlUxhSs84ziIqCAESJgYMqgEgfRvh8Dp4twQb4atKiTLih8ERLofVSBKl4pOrITaPmrsgIioIARImCBSqASBGsUWoN/z0s6B439BLE0gomotHHdPD8VgOTvFDULtQpSA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFG8hbgKzekTtRBgxnwg9q+NEAPkz",
					"data": "Ct8BChsDYWR1bQAUbyFuArN6RO1EGDGfCD2r40QA+TMSAQAaCwgBGAEgASoDAAICIiwIARIFAgSoASAaISBn5IiYMBtU8uy81myDYITtyjoI48Z/GeF4g2T9C8fjdiIqCAESJgQGqAEgATZzD1H0y90hfyQPJfVpzFbKtp203llzhaL2bphVN2wgIioIARImBgyqASB9G+HwOni3BBvhq0qJMuKHwREuh9VIEqXik6shNo+auyAiKggBEiYIFKoBIEaxRag3/PSzoHjf0EsTSCiai0cd08PxWA5O8UNQu1ClIA=="
				  },
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFHQgVx8nLNdp+jNVP71j3Px1xR1Y",
					"data": "Ct4BChsDYWR1bQAUdCBXHycs12n6M1U/vWPc/HXFHVgSAQAaDAgBGAEgASoEAAKoASIqCAESJgIEqAEgdi5sUlGmzJ1LC4RqYwod8LELUa1/LKoq6IGqFIPScoAgIioIARImBAaoASABNnMPUfTL3SF/JA8l9WnMVsq2nbTeWXOFovZumFU3bCAiKggBEiYGDKoBIH0b4fA6eLcEG+GrSoky4ofBES6H1UgSpeKTqyE2j5q7ICIqCAESJggUqgEgRrFFqDf89LOgeN/QSxNIKJqLRx3Tw/FYDk7xQ1C7UKUg"
				  }
				]
			  },
			  {
				"key": "AhQ6yUJFGaPpSdxxcvT9gjIU+nHGlWFkdW0=",
				"value": "Mzk5OTk5OTk5OTk5OTk5OTk5OTk5NzAwMDAw",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "AhRvIW4Cs3pE7UQYMZ8IPavjRAD5M2FkdW0=",
					"data": "CvwBChoCFG8hbgKzekTtRBgxnwg9q+NEAPkzYWR1bRIbMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGgsIARgBIAEqAwACAiIsCAESBQIEqgEgGiEgTKCBvg0DQmIZjd1G9X0/z1PqesM6xLasc6hssDm9fFMiLAgBEgUEBqoBIBohIB++wMEwc7l/SYUEk+aRgmW9SNBg7TuLfMwGXVzo1SbzIiwIARIFBgyqASAaISBk0r+ufEOmlWtaxsEWckD1gPh2kHL0MdheTXmPAQRPwyIqCAESJggUqgEgRrFFqDf89LOgeN/QSxNIKJqLRx3Tw/FYDk7xQ1C7UKUg"
				  },
				  {
					"type": "ics23:iavl",
					"key": "AhQ6yUJFGaPpSdxxcvT9gjIU+nHGlWFkdW0=",
					"data": "Cs8BChoCFDrJQkUZo+lJ3HFy9P2CMhT6ccaVYWR1bRIbMzk5OTk5OTk5OTk5OTk5OTk5OTk5ODAwMDAwGgwIARgBIAEqBAACqgEiLAgBEgUCBKoBIBohIKU8yTwbJN1z9LAsRHsw6vPagW2WjSIdfU1GFPJvQA8wIioIARImBAiqASBs0IEDDkQuoiP16tHyjeQyYCy7GQrh5AwHMr9u0VqOKyAiLAgBEgUIFKoBIBohIJS2u0QtVxEwt9cJS1yaH2d+4cnKUL1eijyfZSFmxajX"
				  },
				  {
					"type": "ics23:iavl",
					"key": "AhRP6nZCe4NFhh6Ao1QKip2Tb9OTkWFkdW0=",
					"data": "CswBChoCFE/qdkJ7g0WGHoCjVAqKnZNv05ORYWR1bRIbNTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGgsIARgBIAEqAwACAiIqCAESJgIEqgEgBKs6C4rEQg4/tq3p4VCuTsYQy+WsB8oiswuUasLj1lEgIioIARImBAiqASBs0IEDDkQuoiP16tHyjeQyYCy7GQrh5AwHMr9u0VqOKyAiLAgBEgUIFKoBIBohIJS2u0QtVxEwt9cJS1yaH2d+4cnKUL1eijyfZSFmxajX"
				  }
				]
			  },
			  {
				"key": "AhR0IFcfJyzXafozVT+9Y9z8dcUdWGFkdW0=",
				"value": "MzAwMDAw",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFDrJQkUZo+lJ3HFy9P2CMhT6ccaV",
					"data": "CrMBChsDYWR1bQAUOslCRRmj6UnccXL0/YIyFPpxxpUSAQAaCwgBGAEgASoDAAICIioIARImBAaqASDIHeINewrAgcTbwOGzbDiNdABmdksZt97BfMrlrspPEiAiLAgBEgUGDKoBIBohIGTSv658Q6aVa1rGwRZyQPWA+HaQcvQx2F5NeY8BBE/DIioIARImCBSsASCvRowGJSk8gAwf4zRJdrKlCUaUPQpm8kdvsSQUb/FF7SA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AhR0IFcfJyzXafozVT+9Y9z8dcUdWGFkdW0=",
					"data": "CuYBChoCFHQgVx8nLNdp+jNVP71j3Px1xR1YYWR1bRIGMjAwMDAwGgwIARgBIAEqBAACqgEiKggBEiYCBKoBIFrzkvkooK9Y1DelQr/gZmzgOa8D0tP2w0CnbIZWHJfvICIsCAESBQQGqgEgGiEgH77AwTBzuX9JhQST5pGCZb1I0GDtO4t8zAZdXOjVJvMiLAgBEgUGDKoBIBohIGTSv658Q6aVa1rGwRZyQPWA+HaQcvQx2F5NeY8BBE/DIioIARImCBSsASCvRowGJSk8gAwf4zRJdrKlCUaUPQpm8kdvsSQUb/FF7SA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "AhRvIW4Cs3pE7UQYMZ8IPavjRAD5M2FkdW0=",
					"data": "CvwBChoCFG8hbgKzekTtRBgxnwg9q+NEAPkzYWR1bRIbMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGgsIARgBIAEqAwACAiIsCAESBQIEqgEgGiEgTKCBvg0DQmIZjd1G9X0/z1PqesM6xLasc6hssDm9fFMiLAgBEgUEBqoBIBohIB++wMEwc7l/SYUEk+aRgmW9SNBg7TuLfMwGXVzo1SbzIiwIARIFBgyqASAaISBk0r+ufEOmlWtaxsEWckD1gPh2kHL0MdheTXmPAQRPwyIqCAESJggUrAEgr0aMBiUpPIAMH+M0SXaypQlGlD0KZvJHb7EkFG/xRe0g"
				  },
				  {
					"type": "ics23:iavl",
					"key": "A2FkdW0AFE/qdkJ7g0WGHoCjVAqKnZNv05OR",
					"data": "CrMBChsDYWR1bQAUT+p2QnuDRYYegKNUCoqdk2/Tk5ESAQAaCwgBGAEgASoDAAICIiwIARIFBAaoASAaISBP3asyRmeyidQGCOiCsJrrkx+sMJGLG40hlUxhSs84ziIqCAESJgYMqgEgfRvh8Dp4twQb4atKiTLih8ERLofVSBKl4pOrITaPmrsgIioIARImCBSsASCvRowGJSk8gAwf4zRJdrKlCUaUPQpm8kdvsSQUb/FF7SA="
				  }
				]
			  }
			]
		  },
		  "capability": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "Y2FwYWJpbGl0eQ==",
			  "data": "CoQCCgpjYXBhYmlsaXR5EiA3lkbhshsNOWB5ZbCtI3H4PuziietGnQJN2P/Dz7sM0BoJCAEYASABKgEAIicIARIBARog7H+dUS6KGVaKZQ88zi+kEDhcuWC0GDk54KcKJP6lbt0iJQgBEiEBwjP3cNfYD1m7NX7IeZclrRmz6IqG1LyPf6aSwPyZoUwiJwgBEgEBGiDZ9FgUGI9NCzQnlU37buiGA2/vUJm82SMz/z1qcNFyfiInCAESAQEaIB51WBI53NBm99ZbtG0TuOZ9wHqDbkgCrq8GtyosjOMWIicIARIBARogmfk2bhwD7Rkn+Vijwi+xup9YpZytzlIISUNMvhS/yxg="
			},
			"root_hash": "N5ZG4bIbDTlgeWWwrSNx+D7s4onrRp0CTdj/w8+7DNA="
		  },
		  "claims": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "Y2xhaW1z",
			  "data": "Cv4BCgZjbGFpbXMSIMwipkZ11C2jPVKIuFIP2m9rivi1vWFbuIhQgseFc9mgGgkIARgBIAEqAQAiJQgBEiEBL+QThvg5Xt8E1S+U6MBmQXPQzo6gHUfh+C3TB9brjgYiJQgBEiEBwjP3cNfYD1m7NX7IeZclrRmz6IqG1LyPf6aSwPyZoUwiJwgBEgEBGiDZ9FgUGI9NCzQnlU37buiGA2/vUJm82SMz/z1qcNFyfiInCAESAQEaIB51WBI53NBm99ZbtG0TuOZ9wHqDbkgCrq8GtyosjOMWIicIARIBARogmfk2bhwD7Rkn+Vijwi+xup9YpZytzlIISUNMvhS/yxg="
			},
			"root_hash": "zCKmRnXULaM9Uoi4Ug/ab2uK+LW9YVu4iFCCx4Vz2aA="
		  },
		  "distribution": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "ZGlzdHJpYnV0aW9u",
			  "data": "CoYCCgxkaXN0cmlidXRpb24SIO5nT55LR93Y86MReqmhh7GVs4GE2sn++6wTb3Vmq2d/GgkIARgBIAEqAQAiJwgBEgEBGiAEOBSG1EaAWmuNUz0i8QumD+3+hen0Rt6kKQqfadTvMyInCAESAQEaIPd88klzfn8XB5UBlKbnNauUuils/GeFDKiDSgc481kWIiUIARIhATDxVA/RtK8YAtXLVY0CMNjiJKMaZ7duYPXL5iDechlMIicIARIBARogHnVYEjnc0Gb31lu0bRO45n3AeoNuSAKurwa3KiyM4xYiJwgBEgEBGiCZ+TZuHAPtGSf5WKPCL7G6n1ilnK3OUghJQ0y+FL/LGA=="
			},
			"root_hash": "7mdPnktH3djzoxF6qaGHsZWzgYTayf77rBNvdWarZ38="
		  },
		  "epochs": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "ZXBvY2hz",
			  "data": "Cv4BCgZlcG9jaHMSIK/rwLOEtMqWvH1nLFjBKdfyE2cF/bU2rw4YJ+liEZ+lGgkIARgBIAEqAQAiJQgBEiEB0KRElae8rS0ZVjidAFUa/v4WTp2qf/SEyruTxmfEJVMiJwgBEgEBGiD3fPJJc35/FweVAZSm5zWrlLopbPxnhQyog0oHOPNZFiIlCAESIQEw8VQP0bSvGALVy1WNAjDY4iSjGme3bmD1y+Yg3nIZTCInCAESAQEaIB51WBI53NBm99ZbtG0TuOZ9wHqDbkgCrq8GtyosjOMWIicIARIBARogmfk2bhwD7Rkn+Vijwi+xup9YpZytzlIISUNMvhS/yxg="
			},
			"root_hash": "r+vAs4S0ypa8fWcsWMEp1/ITZwX9tTavDhgn6WIRn6U="
		  },
		  "erc20": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "ZXJjMjA=",
			  "data": "Cv0BCgVlcmMyMBIghOK4z2sXexNbrFWGNhr04F1zal1IoNokU3tBEeBjdT4aCQgBGAEgASoBACInCAESAQEaIGsyqRCZsPX4UVVvsEmK7QzEkM6a2PgmkPiY7ieRGwTFIiUIARIhAYhq7rUbXffMw/+13tI3PDwKzu8yLMdCuRJRJpFXnnT6IiUIARIhATDxVA/RtK8YAtXLVY0CMNjiJKMaZ7duYPXL5iDechlMIicIARIBARogHnVYEjnc0Gb31lu0bRO45n3AeoNuSAKurwa3KiyM4xYiJwgBEgEBGiCZ+TZuHAPtGSf5WKPCL7G6n1ilnK3OUghJQ0y+FL/LGA=="
			},
			"root_hash": "hOK4z2sXexNbrFWGNhr04F1zal1IoNokU3tBEeBjdT4="
		  },
		  "evm": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "ZXZt",
			  "data": "CvkBCgNldm0SIKwHk0SIb1WZ7UFglE55/2pS32k6BkQWm7ZtL/0kcz73GgkIARgBIAEqAQAiJQgBEiEBs6Un3w1k8F/TxIubmxsDH/NZzHh3ObWtyHYB63gyveUiJQgBEiEBiGrutRtd98zD/7Xe0jc8PArO7zIsx0K5ElEmkVeedPoiJQgBEiEBMPFUD9G0rxgC1ctVjQIw2OIkoxpnt25g9cvmIN5yGUwiJwgBEgEBGiAedVgSOdzQZvfWW7RtE7jmfcB6g25IAq6vBrcqLIzjFiInCAESAQEaIJn5Nm4cA+0ZJ/lYo8IvsbqfWKWcrc5SCElDTL4Uv8sY"
			},
			"root_hash": "rAeTRIhvVZntQWCUTnn/alLfaToGRBabtm0v/SRzPvc="
		  },
		  "feemarket": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "ZmVlbWFya2V0",
			  "data": "CoMCCglmZWVtYXJrZXQSIGW4vJwiDF1XlXAUkeHUnFX7uj7FwsuQ7+K4usuSEEj4GgkIARgBIAEqAQAiJwgBEgEBGiBGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyInCAESAQEaIDm991pzXmJKmUGgm77VMIXsKuWdCDRdDkoO4S9FIIsoIicIARIBARogjjrRqBB1IBFgrOMoZyO+zB1HYakApW4wBnt3otG3SIoiJQgBEiEB8eDdPrfdIBG59Wy7MtCdqL0FZG4A/S2sAq7QkTry7gIiJwgBEgEBGiCZ+TZuHAPtGSf5WKPCL7G6n1ilnK3OUghJQ0y+FL/LGA=="
			},
			"root_hash": "Zbi8nCIMXVeVcBSR4dScVfu6PsXCy5Dv4ri6y5IQSPg="
		  },
		  "gov": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "Z292",
			  "data": "CvsBCgNnb3YSIGDNE3wZYuysYWOJ1oA0gz8pIVCcLShapfUVOZfOlop0GgkIARgBIAEqAQAiJQgBEiEB9JdcB+c6im8p3ikfWhyn1zQ3Utit5Iik6T3KlOwcpg4iJwgBEgEBGiA5vfdac15iSplBoJu+1TCF7CrlnQg0XQ5KDuEvRSCLKCInCAESAQEaII460agQdSARYKzjKGcjvswdR2GpAKVuMAZ7d6LRt0iKIiUIARIhAfHg3T633SARufVsuzLQnai9BWRuAP0trAKu0JE68u4CIicIARIBARogmfk2bhwD7Rkn+Vijwi+xup9YpZytzlIISUNMvhS/yxg="
			},
			"root_hash": "YM0TfBli7KxhY4nWgDSDPykhUJwtKFql9RU5l86WinQ="
		  },
		  "ibc": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "aWJj",
			  "data": "CvsBCgNpYmMSIMYX7vgtNQhZ/EwbARgHn60g7NHl/iKUyGqG6UwNFMpVGgkIARgBIAEqAQAiJwgBEgEBGiAT2oLbao1mj7fNjPLv55ngBccbnj1P0bPzxy3Sb4nNNiIlCAESIQH/F4vWMCz0GPlERAFl9JtzTFT0lOB5PyIS2pcrQ5bqiyInCAESAQEaII460agQdSARYKzjKGcjvswdR2GpAKVuMAZ7d6LRt0iKIiUIARIhAfHg3T633SARufVsuzLQnai9BWRuAP0trAKu0JE68u4CIicIARIBARogmfk2bhwD7Rkn+Vijwi+xup9YpZytzlIISUNMvhS/yxg="
			},
			"root_hash": "xhfu+C01CFn8TBsBGAefrSDs0eX+IpTIaobpTA0UylU="
		  },
		  "mint": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "bWludA==",
			  "data": "CvoBCgRtaW50EiDqC4Aa2VZzlcBpmPRg4F1MOzWJ+E5jxtU3VC3Em+6M+RoJCAEYASABKgEAIiUIARIhAQpu6OgOMgYPQCwhx30za2BMuXJQ8+Z2CMWcaW2bnt7hIiUIARIhAf8Xi9YwLPQY+UREAWX0m3NMVPSU4Hk/IhLalytDluqLIicIARIBARogjjrRqBB1IBFgrOMoZyO+zB1HYakApW4wBnt3otG3SIoiJQgBEiEB8eDdPrfdIBG59Wy7MtCdqL0FZG4A/S2sAq7QkTry7gIiJwgBEgEBGiCZ+TZuHAPtGSf5WKPCL7G6n1ilnK3OUghJQ0y+FL/LGA=="
			},
			"root_hash": "6guAGtlWc5XAaZj0YOBdTDs1ifhOY8bVN1QtxJvujPk="
		  },
		  "params": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "cGFyYW1z",
			  "data": "Cv4BCgZwYXJhbXMSINaOg2YbVTnus/mnXB8CCHspuXQb+n1xOnzWo9C4zpjIGgkIARgBIAEqAQAiJwgBEgEBGiDaYcZ8/Dcx0cQUiMg8Q1qYsDohW/UDGGOt3acZEg4UyCInCAESAQEaIM/Anlr0MvPS44pjriUxqKCgl4mOtI+7unGF+E5rv5bUIiUIARIhAVAKZQvpYbchIznWsjw/j+rgCcbocN9mFp3uL93oDBdfIiUIARIhAfHg3T633SARufVsuzLQnai9BWRuAP0trAKu0JE68u4CIicIARIBARogmfk2bhwD7Rkn+Vijwi+xup9YpZytzlIISUNMvhS/yxg="
			},
			"root_hash": "1o6DZhtVOe6z+adcHwIIeym5dBv6fXE6fNaj0LjOmMg=",
			"witness_data": [
			  {
				"operation": 1,
				"key": "YmFuay9TZW5kRW5hYmxlZA==",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "bWludC9FcG9jaElkZW50aWZpZXI=",
					"data": "CooCChRtaW50L0Vwb2NoSWRlbnRpZmllchIGIm1pbnQiGgsIARgBIAEqAwACAiIrCAESBAIEAiAaISD0Pg3VYxCSprZSmG0aKo7OmqrFsWg/29BTExoXBnlvfiIrCAESBAQIAiAaISCsOUkc5BxAX4eLVb2dRQSkdCkNGVVLqJwshI9H2SUIFCIrCAESBAYQAiAaISBKXHxChaX4Ep5Qt0pasgMEG5Y4aXhljU+QOOhMw3umkCIrCAESBAoiAiAaISBsCtBFdTusrZu0cUtURAoqdZ2GyOu5dpte40ZpPvhloCIpCAESJQxCAiApIRNaHCH7B9WEywgDp37O7ibUeMa6i1cZHQJcZT8ZnCA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "ZGlzdHJpYnV0aW9uL2JvbnVzcHJvcG9zZXJyZXdhcmQ=",
					"data": "CqYCCiBkaXN0cmlidXRpb24vYm9udXNwcm9wb3NlcnJld2FyZBIWIjAuMDQwMDAwMDAwMDAwMDAwMDAwIhoLCAEYASABKgMAAgIiKwgBEgQCBAIgGiEgaEi7qdP26XWmzXX160QUJ+JFOUakM4Qp32Lz2Uk72X4iKwgBEgQECAIgGiEgADSgUws1tSqnxn8w7ZeqU33+gYA8vAJ4BpuKGCO/IOIiKwgBEgQGEAIgGiEgRf0azEtCUjo8aX6aZM5ydTAVr+VhekPILcM7KqqK6ZwiKQgBEiUIIAIgPqU+zbDYyx/Kj/Bs9xzyUwoCSMJe910rYsKMmBB2kc0gIisIARIEDEICIBohINyEzrk8ev7lT8c9Jbro7aAsWiiiuVGam3UIXxVLwAJz"
				  },
				  {
					"type": "ics23:iavl",
					"key": "YXV0aC9UeFNpemVDb3N0UGVyQnl0ZQ==",
					"data": "CooCChZhdXRoL1R4U2l6ZUNvc3RQZXJCeXRlEgQiMTAiGgsIARgBIAEqAwACAiIrCAESBAIEAiAaISD23HVFgbNeLuRPDz8DZHufddFjCuoP1szIGO5sU44uviIrCAESBAQIAiAaISDEz3yuhYxQqGbq+h4VYLy+wYilnXl89XNtb/+kOPjKCCIpCAESJQYQAiDawnOG+RCzaRndZZA3za+wdtknQoT0hSfDVzPPV+hfQCAiKwgBEgQIIAIgGiEgMT/9QcLd8OSoXoW298Bxc0qKOTXkMrPJO2z8NWibyHgiKwgBEgQMQgIgGiEg3ITOuTx6/uVPxz0luujtoCxaKKK5UZqbdQhfFUvAAnM="
				  },
				  {
					"type": "ics23:iavl",
					"key": "YmFuay9TZW5kRW5hYmxlZA==",
					"data": "CoACChBiYW5rL1NlbmRFbmFibGVkEgJbXRoLCAEYASABKgMAAgIiKwgBEgQCBAIgGiEghG4EvuU1t2BOlWH24h7bfnOysW36Dg/AdNZpPZsQBYUiKQgBEiUECAIg4J34dp/LO1Ar0oeiHgbtKPwBZ46Gd/6OL/iu4cwGJ+sgIikIARIlBhACINrCc4b5ELNpGd1lkDfNr7B22SdChPSFJ8NXM89X6F9AICIrCAESBAggAiAaISAxP/1Bwt3w5Khehbb3wHFzSoo5NeQys8k7bPw1aJvIeCIrCAESBAxCAiAaISDchM65PHr+5U/HPSW66O2gLFooorlRmpt1CF8VS8ACcw=="
				  },
				  {
					"type": "ics23:iavl",
					"key": "ZGlzdHJpYnV0aW9uL2Jhc2Vwcm9wb3NlcnJld2FyZA==",
					"data": "CqECCh9kaXN0cmlidXRpb24vYmFzZXByb3Bvc2VycmV3YXJkEhYiMC44MDAwMDAwMDAwMDAwMDAwMDAiGgsIARgBIAEqAwACAiIpCAESJQIEAiDMy9wXNLS7GcnKMe4e9d4Y/FVdONww4OLL6j/YeCR24iAiKQgBEiUECAIg4J34dp/LO1Ar0oeiHgbtKPwBZ46Gd/6OL/iu4cwGJ+sgIikIARIlBhACINrCc4b5ELNpGd1lkDfNr7B22SdChPSFJ8NXM89X6F9AICIrCAESBAggAiAaISAxP/1Bwt3w5Khehbb3wHFzSoo5NeQys8k7bPw1aJvIeCIrCAESBAxCAiAaISDchM65PHr+5U/HPSW66O2gLFooorlRmpt1CF8VS8ACcw=="
				  }
				]
			  },
			  {
				"operation": 1,
				"key": "YmFuay9EZWZhdWx0U2VuZEVuYWJsZWQ=",
				"proofs": [
				  {
					"type": "ics23:iavl",
					"key": "bWludC9FcG9jaElkZW50aWZpZXI=",
					"data": "CooCChRtaW50L0Vwb2NoSWRlbnRpZmllchIGIm1pbnQiGgsIARgBIAEqAwACAiIrCAESBAIEAiAaISD0Pg3VYxCSprZSmG0aKo7OmqrFsWg/29BTExoXBnlvfiIrCAESBAQIAiAaISCsOUkc5BxAX4eLVb2dRQSkdCkNGVVLqJwshI9H2SUIFCIrCAESBAYQAiAaISBKXHxChaX4Ep5Qt0pasgMEG5Y4aXhljU+QOOhMw3umkCIrCAESBAoiAiAaISBsCtBFdTusrZu0cUtURAoqdZ2GyOu5dpte40ZpPvhloCIpCAESJQxCAiApIRNaHCH7B9WEywgDp37O7ibUeMa6i1cZHQJcZT8ZnCA="
				  },
				  {
					"type": "ics23:iavl",
					"key": "ZGlzdHJpYnV0aW9uL2JvbnVzcHJvcG9zZXJyZXdhcmQ=",
					"data": "CqYCCiBkaXN0cmlidXRpb24vYm9udXNwcm9wb3NlcnJld2FyZBIWIjAuMDQwMDAwMDAwMDAwMDAwMDAwIhoLCAEYASABKgMAAgIiKwgBEgQCBAIgGiEgaEi7qdP26XWmzXX160QUJ+JFOUakM4Qp32Lz2Uk72X4iKwgBEgQECAIgGiEgADSgUws1tSqnxn8w7ZeqU33+gYA8vAJ4BpuKGCO/IOIiKwgBEgQGEAIgGiEgRf0azEtCUjo8aX6aZM5ydTAVr+VhekPILcM7KqqK6ZwiKQgBEiUIIAIgPqU+zbDYyx/Kj/Bs9xzyUwoCSMJe910rYsKMmBB2kc0gIisIARIEDEICIBohINyEzrk8ev7lT8c9Jbro7aAsWiiiuVGam3UIXxVLwAJz"
				  },
				  {
					"type": "ics23:iavl",
					"key": "YXV0aC9UeFNpemVDb3N0UGVyQnl0ZQ==",
					"data": "CooCChZhdXRoL1R4U2l6ZUNvc3RQZXJCeXRlEgQiMTAiGgsIARgBIAEqAwACAiIrCAESBAIEAiAaISD23HVFgbNeLuRPDz8DZHufddFjCuoP1szIGO5sU44uviIrCAESBAQIAiAaISDEz3yuhYxQqGbq+h4VYLy+wYilnXl89XNtb/+kOPjKCCIpCAESJQYQAiDawnOG+RCzaRndZZA3za+wdtknQoT0hSfDVzPPV+hfQCAiKwgBEgQIIAIgGiEgMT/9QcLd8OSoXoW298Bxc0qKOTXkMrPJO2z8NWibyHgiKwgBEgQMQgIgGiEg3ITOuTx6/uVPxz0luujtoCxaKKK5UZqbdQhfFUvAAnM="
				  },
				  {
					"type": "ics23:iavl",
					"key": "YmFuay9TZW5kRW5hYmxlZA==",
					"data": "CoACChBiYW5rL1NlbmRFbmFibGVkEgJbXRoLCAEYASABKgMAAgIiKwgBEgQCBAIgGiEghG4EvuU1t2BOlWH24h7bfnOysW36Dg/AdNZpPZsQBYUiKQgBEiUECAIg4J34dp/LO1Ar0oeiHgbtKPwBZ46Gd/6OL/iu4cwGJ+sgIikIARIlBhACINrCc4b5ELNpGd1lkDfNr7B22SdChPSFJ8NXM89X6F9AICIrCAESBAggAiAaISAxP/1Bwt3w5Khehbb3wHFzSoo5NeQys8k7bPw1aJvIeCIrCAESBAxCAiAaISDchM65PHr+5U/HPSW66O2gLFooorlRmpt1CF8VS8ACcw=="
				  },
				  {
					"type": "ics23:iavl",
					"key": "YmFuay9EZWZhdWx0U2VuZEVuYWJsZWQ=",
					"data": "CokCChdiYW5rL0RlZmF1bHRTZW5kRW5hYmxlZBIEdHJ1ZRoLCAEYASABKgMAAgIiKQgBEiUCBAIgUOvhTILubHeVEjASYPHz2/ObEiwy3i8ebwTQUmJokFQgIisIARIEBAgCIBohIMTPfK6FjFCoZur6HhVgvL7BiKWdeXz1c21v/6Q4+MoIIikIARIlBhACINrCc4b5ELNpGd1lkDfNr7B22SdChPSFJ8NXM89X6F9AICIrCAESBAggAiAaISAxP/1Bwt3w5Khehbb3wHFzSoo5NeQys8k7bPw1aJvIeCIrCAESBAxCAiAaISDchM65PHr+5U/HPSW66O2gLFooorlRmpt1CF8VS8ACcw=="
				  }
				]
			  }
			]
		  },
		  "sequencers": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "c2VxdWVuY2Vycw==",
			  "data": "CoACCgpzZXF1ZW5jZXJzEiDlM/pUG5TMKKp7y38l+3mlmc3/U+Gj7Y3mdkaGTIyi6xoJCAEYASABKgEAIiUIARIhAZtmtBnHA+2rxeAaRD36f0b1GdQcSbPfBJOWBU3LBCZMIicIARIBARogz8CeWvQy89LjimOuJTGooKCXiY60j7u6cYX4Tmu/ltQiJQgBEiEBUAplC+lhtyEjOdayPD+P6uAJxuhw32YWne4v3egMF18iJQgBEiEB8eDdPrfdIBG59Wy7MtCdqL0FZG4A/S2sAq7QkTry7gIiJwgBEgEBGiCZ+TZuHAPtGSf5WKPCL7G6n1ilnK3OUghJQ0y+FL/LGA=="
			},
			"root_hash": "5TP6VBuUzCiqe8t/Jft5pZnN/1Pho+2N5nZGhkyMous="
		  },
		  "staking": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "c3Rha2luZw==",
			  "data": "Cv0BCgdzdGFraW5nEiDvv1OaAqXPu/ctxYjvsMHUTkHzW4QI9MUcnc4HCQuqdhoJCAEYASABKgEAIicIARIBARogOqsHULjzmZkig3Kxczq2JoCMuiq6iXWpKHea7ZB9gWAiJQgBEiEBnXgusslqDIXWvhWQlB1uPZ9JIiTT7rqo2Kxqv0sz1JsiJQgBEiEBUAplC+lhtyEjOdayPD+P6uAJxuhw32YWne4v3egMF18iJQgBEiEB8eDdPrfdIBG59Wy7MtCdqL0FZG4A/S2sAq7QkTry7gIiJwgBEgEBGiCZ+TZuHAPtGSf5WKPCL7G6n1ilnK3OUghJQ0y+FL/LGA=="
			},
			"root_hash": "779TmgKlz7v3LcWI77DB1E5B81uECPTFHJ3OBwkLqnY="
		  },
		  "transfer": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "dHJhbnNmZXI=",
			  "data": "CvwBCgh0cmFuc2ZlchIgJZrOM/PYGf2cx1GWjMg9UDqS6Kfm/V4ymAaCSVZf1wYaCQgBGAEgASoBACIlCAESIQESU8NHn8rclOVLUR9yPt1Yq8shWX5Q23XP3rbHPiz6XSIlCAESIQGdeC6yyWoMhda+FZCUHW49n0kiJNPuuqjYrGq/SzPUmyIlCAESIQFQCmUL6WG3ISM51rI8P4/q4AnG6HDfZhad7i/d6AwXXyIlCAESIQHx4N0+t90gEbn1bLsy0J2ovQVkbgD9LawCrtCROvLuAiInCAESAQEaIJn5Nm4cA+0ZJ/lYo8IvsbqfWKWcrc5SCElDTL4Uv8sY"
			},
			"root_hash": "JZrOM/PYGf2cx1GWjMg9UDqS6Kfm/V4ymAaCSVZf1wY="
		  },
		  "upgrade": {
			"proof": {
			  "type": "ics23:simple",
			  "key": "dXBncmFkZQ==",
			  "data": "Cl0KB3VwZ3JhZGUSIBOxv83jb6bUEnz6uY+teBuvQAGPTZkAfmGg48A+BVc9GgkIARgBIAEqAQAiJQgBEiEB6QN0fdqEtGjpV6gY0rUjtFsCrx8TvcGIZZgz1iaRBrc="
			},
			"root_hash": "E7G/zeNvptQSfPq5j614G69AAY9NmQB+YaDjwD4FVz0="
		  }
		},
		"fraudulent_deliver_tx": {
		  "tx": "Co0BCooBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEmoKK2V0aG0xOHR5NXkzZ2U1MDU1bmhyM3d0NjBtcTNqem5hOHIzNTRlNGo4ZmMSK2V0aG0xd3NzOXc4ZTg5bnRrbjczbjI1bG02Yzd1bDM2dTI4MmM0c3E1cW0aDgoEYWR1bRIGMTAwMDAwEmEKWQpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQJn8mJfKSWDNyaY7onyhXJN3ZVUVxs6TSxX7j4/gIuoQRIECgIIARgDEgQQwJoMGkF7GTUWou/fCk4TPjTFCX8S+vbPL30S7LE2t9XG3SRoDCWTolBWcRaUlyZwqTLx66Farm3KZmRCeaT6w0O9OEklAA=="
		}
	  }`
)
