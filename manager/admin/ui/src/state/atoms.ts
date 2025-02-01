import { atom } from "recoil";
import { Service, TableInfo } from "./models";

export const loadingState = atom<boolean>({
  key: "Loading",
  default: false,
});

export const tableListState = atom<TableInfo[]>({
  key: "TableList",
  default: [],
});

export const serviceListState = atom<Service[]>({
  key: "ServiceList",
  default: [],
});
