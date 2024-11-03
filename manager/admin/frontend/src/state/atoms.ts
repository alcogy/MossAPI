import { atom } from "recoil";
import { dmServiceList, dmTableList, Service, Table } from "./models";

export const tableListState = atom<Table[]>({
  key: "TableList",
  default: dmTableList,
});

export const serviceListState = atom<Service[]>({
  key: "ServiceList",
  default: dmServiceList,
});
