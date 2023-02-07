package json

type Data struct {
	InterestRate   float64 `json:"InterestRate" binding:"required"`
	NumberOfYears int64   `json:"NumberOfYears" binding:"required"`
	FutureValue   int64   `json:"FutureValue" binding:"required"`
	PresentValue  int64   `json:"PresentValue" binding:"required"`
}

type PublicData struct {
	InterestRate   float64 `json:"InterestRate" binding:"required"`
	NumberOfYears int64   `json:"NumberOfYears" binding:"required"`
	Proof         string  `json:"Proof" binding:"required"`
	VerifyingKey  string  `json:"VerifyingKey" binding:"required"`
}

type ProveResult struct {
	Proof        string `json:"Proof" example:"a253e76c29eee51bdb267fd8d62685dd5d403f83b1d3bf4f2107a177339b9a7195861c57ea7508f77537779907c0bf7ed2fb1ae82a1933b841153bfc54365f5025181c033e24b0700f756ea96d72781de8a99858865b05cab03669797d65dd9d92d9a9596fee5feafabd760a2fa7b28717735ccb52f4350695ee4bcad4a783ca"`
	VerifyingKey string `json:"VerifyingKey" example:"9c3930900244e5c02daba254987d7884d8379f735ee832c906d2d692578923dfc9f03e86084213ec2828d355ec810ec35753b3250b60938b7293bc5cebb33202c04778fd09b4ecbe7e90c2ae0aeecd1ba5e15efa07cc5f4690d6b931b890aca614c08c95fb1ba104357928904454fd463848d9c5a667a7970de87303ae9bec4c90d0308db9276154ea389ab9a09987d57a0a519c1201bf0750fd7bd1278b11992914f8e9e2283978651ec3a960aa8183e706fdb96264881dc90e51f27d39f12ecf33fcdf856c3ff6cb3ffbc48f136c79be6a9e5d7f30d53f32be564363237feb917958eb330b7c9a67c167ad0a0eeaabb71f16d48c6d51e3064d8b57b499bbdd0887359b7bf7ba632f1c7a2441db7a480171149c45c805e16087b9fad0b06fe600000006ee52afb85e03cd8e5aedfeca6e0ab8fc096bb59e854cb6874d7a0c8812d0cf9e833d0069f747db8f7b360c98ccd518248494bb11834734a1a6dbc0c48d5fdc418aab9680d07922668584d4e92e2a1203258d6e3d84622108667f4d728832e392e4fd6fc43e6694972ce184d50f6328ff2eac376c960b2a1a82bd6bb1707cc62acf1d9e279e143ed59a85b3aec7294f61c0cfbefea9632d392d3127f7419bb7f7de5c45973a6c99e0b255873c746000369245c394924408cf1b7b1a97298cc7f6"`
}

type VerifyResult struct {
	Message string `json:"message" example:"Valid Proof"`
}

type BadProveResult struct {
	Message string `json:"message" example:"constraint #XXXX is not satisfied"`
}

type BadVerifyResult struct {
	Message string `json:"message" example:"pairing doesn't match"`
}
