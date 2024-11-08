export interface Table {
  name: string;
  desc: string;
  columns?: Column[];
}

export interface Column {
  name: string;
  type: number;
  size: number;
  pk: boolean;
  notNull: boolean;
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
  | "notNull"
  | "size";

export interface Service {
  id: string;
  name: string;
  port: string;
  status: string;
}
