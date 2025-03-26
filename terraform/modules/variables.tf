variable "app" {
  description = "アプリケーション名"
  type        = string
}

variable "env" {
  description = "環境名"
  type        = string
}

variable "db_username" {
  description = "DB ユーザー名 (手動でユーザーを作成して、権限を GRANT すること)"
  type        = string
}

variable "db_password" {
  description = "DB パスワード (手動でユーザーを作成して、権限を GRANT すること)"
  type        = string
}

variable "db_name" {
  description = "データベース名"
  type        = string
}

variable "db_port" {
  description = "データベースポート"
  type        = number
  default     = 3306
}

variable "db_master_username" {
  description = "DB ユーザー名"
  type        = string
}
variable "db_master_password" {
  description = "DB パスワード"
  type        = string
}

variable "identity_api_base_path" {
  description = "Ludens User Console identity API ベースパス"
  type        = string
}

variable "service" {
  description = "サービス名 (Ludens User Console との通信用)"
  type        = string
}

variable "service_kid" {
  description = "サービス KID (Ludens User Console との通信用)"
  type        = string
}

variable "service_key" {
  description = "サービスキー (Ludens User Console との通信用)"
  type        = string
}

variable "firebase_project_id" {
  description = "Firebase Project ID"
  type        = string
}
