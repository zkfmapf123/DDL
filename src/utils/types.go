package utils

type LambdaTags struct {
	Name                 string `json:"name"`          // 이름
	Version              string `json:"version"`       // 버전
	Ecr                  string `json:"ecr"`           // ECR
	SSMKeys              string `json:"ssm_key"`       // SSM Key Path
	IsContainer          bool   `json:"is_container"`  // Container 유무
	IsAPIGateway         bool   `json:"is_apigateway"` // API Gateway 유무 (추후 개발)
	RoutePath            string `json:"route_path"`    // API Gateway RoutePath (추후 개발)
	CreatedAt            string `json:"created_at"`    // 생성날짜
	UpdatedAt            string `json:"updated_at"`    // 업데이트 / 배포 날짜
	Author               string `json:"author"`        // 생성자
	Least_Author         string `json:"least_author"`  // 최근 수정 자
	Least_Deploy_Message string `json:"least_deploy_message"`
}

type ECRTags struct {
	Description     string `json:"description"`      // 배포 메시지
	PreviousVersion string `json:"previous_version"` // 이전 버전 숫자
	Version         string `json:"version"`          // 현재 버전
	Deployer        string `json:"deployer"`         // 배포자
	Profile         string `json:"profile"`          // 배포 profile
	CreatedAt       string `json:"created_at"`       // 생성일자
}
