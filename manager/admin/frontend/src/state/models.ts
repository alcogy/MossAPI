export interface Table {
  name: string;
}

export interface Service {
  id: string;
  name: string;
  port: string;
  status: string;
}

export const dmTableList: Table[] = [
  { name: "customer" },
  { name: "project" },
  { name: "phase" },
  { name: "history" },
];

export const dmServiceList: Service[] = [
  {
    id: "slx7gtiurkjd",
    name: "customer",
    port: "12001",
    status: "Running",
  },
  {
    id: "ew89fjfllowh",
    name: "project",
    port: "12002",
    status: "Running",
  },
  {
    id: "ghuty47yh3uy",
    name: "price",
    port: "12003",
    status: "Stop",
  },
];
