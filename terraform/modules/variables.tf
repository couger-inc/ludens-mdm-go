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