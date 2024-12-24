export interface TableInfo {
  tableName: string;
  tableDesc: string;
  columns?: Column[];
}

export interface Column {
  name: string;
  type: number;
  size: number;
  pk: boolean;
  nullable: boolean;
  unique: number;
  index: number;
  comment: string;
}

export type ColumnFormParams =
  | "name"
  | "comment"
  | "type"
  | "index"
  | "unique"
  | "pk"
  | "nullable"
  | "size";

export interface Service {
  id: string;
  name: string;
  port: string;
  state: string;
  status: string;
}
