import Typography from "@mui/material/Typography";
import ModuleTitle from "../components/ModuleTitle";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { IconButton, TextField } from "@mui/material";
import { useCallback, useEffect, useState } from "react";
import AddBoxIcon from "@mui/icons-material/AddBox";
import { API_TABLE_CREATE, typeList } from "../common/constants";
import { Column, ColumnFormParams } from "../state/models";
import TableForm from "../components/TableForm";

export const initColunInfo: Column = {
  name: "",
  type: 10,
  size: 0,
  pk: false,
  nullable: false,
  unique: 0,
  index: 0,
  comment: "",
};

export default function CreateTable() {
  const [tableName, setTableName] = useState<string>("");
  const [tableDesc, setTableDesc] = useState<string>("");
  const [columns, setColumns] = useState<Column[]>([{ ...initColunInfo }]);

  const disabled = (): boolean => {
    if (tableName === "") return true;
    if (columns.length === 0) return false;
    for (const col of columns) {
      if (col.name === "") return true;
    }
    // TODO regex only alphabet, num, underscore.

    return false;
  };

  const onClickDeleteRow = (index: number) => {
    if (!window.confirm("Delete this column?")) return;
    setColumns(columns.filter((_, i) => i !== index));
  };

  const updateForm = (index: number, kind: ColumnFormParams, value: any) => {
    const newColumns = [...columns];
    const newState = newColumns[index];
    switch (kind) {
      case "name":
      case "comment":
        newState[kind] = value as string;
        break;
      case "type":
      case "index":
      case "unique":
        newState[kind] = value as number;
        break;
      case "pk":
        const v = value as boolean;
        if (v) {
          newState["nullable"] = false;
        }
        newState["pk"] = v;
        break;
      case "nullable":
        newState[kind] = value as boolean;
        break;
      case "size":
        const num = Number(value);
        if (isNaN(num) || num < 0 || num > 255) return;
        newState[kind] = num;
    }
    setColumns(newColumns);
  };

  const onClickCreate = async () => {
    if (!window.confirm("Do you registrate data?")) return;
    const result = await fetch(API_TABLE_CREATE, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        tableName: tableName,
        tableDesc: tableDesc,
        columns: columns.map((v) => {
          return { ...v, type: getTypeLabel(v) };
        }),
      }),
    });
    const json = await result.json();
    if (json["message"] !== "ok") {
      alert("error occured.");
      console.error(json["message"]);
    }
    setTableName("");
    setColumns([{ ...initColunInfo }]);
  };

  const getTypeLabel = (column: Column) => {
    if (column.type === 20) return "varchar(" + column.size + ")";
    for (const t of typeList) {
      if (t.value === column.type) return t.label;
    }
    return "";
  };

  useEffect(() => {
    return () => {
      setColumns([{ ...initColunInfo }]);
    };
  }, []);

  return (
    <Box>
      <ModuleTitle label="Create Table" />
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: 3,
        }}
      >
        <TextField
          label="Table name"
          variant="outlined"
          sx={{ maxWidth: "320px" }}
          value={tableName}
          onChange={(e) => setTableName(e.target.value)}
        />

        <TextField
          label="Table description"
          variant="outlined"
          sx={{ maxWidth: "720px" }}
          value={tableDesc}
          onChange={(e) => setTableDesc(e.target.value)}
        />

        <Box>
          <Typography
            variant="subtitle1"
            sx={{ fontWeight: "bold", marginBottom: 2 }}
          >
            Setting Columns
          </Typography>
          <Box sx={{ display: "flex", flexDirection: "column", gap: 3 }}>
            {columns.length === 0 && (
              <Typography>Please add column...</Typography>
            )}
            {columns.map((value, index) => (
              <TableForm
                key={index}
                value={value}
                index={index}
                updateForm={updateForm}
                onClickDeleteRow={onClickDeleteRow}
              />
            ))}
          </Box>
          <Box sx={{ marginTop: 1 }}>
            <IconButton
              onClick={() => setColumns([...columns, { ...initColunInfo }])}
            >
              <AddBoxIcon />
            </IconButton>
          </Box>
        </Box>

        <Box sx={{ marginTop: 4, display: "flex", gap: 2 }}>
          <Button
            variant="contained"
            onClick={onClickCreate}
            disabled={disabled()}
          >
            Create
          </Button>
          <Button variant="outlined" color="secondary" href="/">
            Cancel
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
