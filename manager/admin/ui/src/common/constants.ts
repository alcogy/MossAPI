import { Column } from "../state/models";

const urlRoot = "http://localhost:5500/api/";

export const API_GET_INFRASTRUCTURE = urlRoot + "infrastructure";

export const API_GET_SERVICES = urlRoot + "services";
export const API_SERVICE_CREATE = urlRoot + "service/create";
export const API_SERVICE_START = urlRoot + "service/start/";
export const API_SERVICE_STOP = urlRoot + "service/stop/";
export const API_SERVICE_REMOVE = urlRoot + "service/remove/";

export const API_GET_TABLES = urlRoot + "tables";
export const API_GET_TABLE_DETAIL = urlRoot + "table/";
export const API_TABLE_CREATE = urlRoot + "table";
export const API_TABLE_UPDATE = urlRoot + "table";
export const API_TABLE_DELETE = urlRoot + "table/";

export const typeList = [
  { value: 10, label: "int" },
  { value: 11, label: "tinyint" },
  { value: 20, label: "varchar" },
  { value: 23, label: "text" },
  { value: 30, label: "date" },
  { value: 31, label: "time" },
  { value: 32, label: "datetime" },
];

export const segmentList = [
  { value: 0, label: "None" },
  { value: 1, label: "1" },
  { value: 2, label: "2" },
  { value: 3, label: "3" },
  { value: 4, label: "4" },
  { value: 5, label: "5" },
  { value: 6, label: "6" },
  { value: 7, label: "7" },
  { value: 8, label: "8" },
  { value: 9, label: "9" },
  { value: 10, label: "10" },
  { value: 11, label: "11" },
  { value: 12, label: "12" },
  { value: 13, label: "13" },
  { value: 14, label: "14" },
  { value: 15, label: "15" },
  { value: 16, label: "16" },
];

export const INIT_COLUMN_INFO: Column = {
  name: "",
  type: 10,
  size: 0,
  pk: false,
  nullable: false,
  unique: 0,
  index: 0,
  comment: "",
};
