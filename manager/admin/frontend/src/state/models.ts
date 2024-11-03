export interface Table {
  name: string;
  columns?: Column[];
}

export interface Column {
  name: string;
  type: string;
  pk: boolean;
  index: boolean;
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

export const dbTableDetail: Table = {
  name: "customer",
  columns: [
    {
      name: "id",
      type: "int",
      pk: true,
      index: false,
    },
    {
      name: "name",
      type: "varchar(255)",
      pk: false,
      index: false,
    },
    {
      name: "area_id",
      type: "int",
      pk: false,
      index: true,
    },
    {
      name: "created_at",
      type: "datetime",
      pk: false,
      index: false,
    },
  ],
};

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
