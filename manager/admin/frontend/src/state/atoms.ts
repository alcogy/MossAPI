import { atom } from "recoil";
import { Service, Table } from "./models";

export const tableListState = atom<Table[]>({
  key: "TableList",
  default: [],
});

export const serviceListState = atom<Service[]>({
  key: "ServiceList",
  default: [],
});
