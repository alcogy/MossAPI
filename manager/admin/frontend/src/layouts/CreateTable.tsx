import Typography from "@mui/material/Typography";
import ModuleTitle from "../components/ModuleTitle";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import {
  Checkbox,
  FormControl,
  FormControlLabel,
  IconButton,
  InputLabel,
  MenuItem,
  Select,
  TextField,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import { useEffect, useState } from "react";
import AddBoxIcon from "@mui/icons-material/AddBox";
import { API_TABLE_CREATE } from "../common/constants";

export interface ColumnForm {
  name: string;
  type: number;
  pk: boolean;
  index: boolean;
}

const typeList = [
  { value: 10, label: "int" },
  { value: 11, label: "tinyint" },
  { value: 20, label: "varchar(255)" },
  { value: 23, label: "text" },
  { value: 30, label: "date" },
  { value: 31, label: "time" },
  { value: 32, label: "datetime" },
];

const initColunInfo: ColumnForm = {
  name: "",
  type: 10,
  pk: false,
  index: false,
};

export default function CreateTable() {
  const [tableName, setTableName] = useState<string>("");
  const [columns, setColumns] = useState<ColumnForm[]>([{ ...initColunInfo }]);

  const disabled = (): boolean => {
    if (tableName === "") return true;
    if (columns.length === 0) return false;
    for (const col of columns) {
      if (col.name === "") return true;
    }
    // TODO regex only alphabet, num, underscore.

    return false;
  };

  const updateForm = (index: number, kind: string, value: any) => {
    const newColumns = [...columns];
    const newState = newColumns[index];
    switch (kind) {
      case "name":
        newState[kind] = value as string;
        break;
      case "type":
        newState[kind] = value as number;
        break;
      case "pk":
        newState[kind] = value as boolean;
        break;
      case "index":
        newState[kind] = value as boolean;
        break;
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
        name: tableName,
        columns: columns.map((v) => {
          return { ...v, type: getTypeLabel(v.type) };
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

  const getTypeLabel = (type: number) => {
    for (const t of typeList) {
      if (t.value === type) return t.label;
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

        <Box>
          <Box
            sx={{
              display: "flex",
              gap: 8,
              alignItems: "center",
              marginBottom: 3,
            }}
          >
            <Typography variant="subtitle1" sx={{ fontWeight: "bold" }}>
              Setting Columns
            </Typography>
            <IconButton
              onClick={() => setColumns([...columns, { ...initColunInfo }])}
            >
              <AddBoxIcon />
            </IconButton>
          </Box>
          <Box sx={{ display: "flex", flexDirection: "column", gap: 3 }}>
            {columns.length === 0 && (
              <Typography>Please add column...</Typography>
            )}
            {columns.map((value, index) => (
              <Box sx={{ display: "flex", gap: 1 }} key={index}>
                <TextField
                  label="Column name"
                  variant="outlined"
                  value={value.name}
                  onChange={(e) => updateForm(index, "name", e.target.value)}
                />
                <FormControl sx={{ minWidth: "148px" }}>
                  <InputLabel id="select-label">Column Type</InputLabel>
                  <Select
                    label="Column Type"
                    labelId="select-label"
                    defaultValue={10}
                    value={value.type}
                    onChange={(e) => updateForm(index, "type", e.target.value)}
                  >
                    {typeList.map((v) => (
                      <MenuItem value={v.value} key={v.value}>
                        {v.label}
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>
                <FormControlLabel
                  control={
                    <Checkbox
                      checked={value.pk}
                      onChange={(e) =>
                        updateForm(index, "pk", e.target.checked)
                      }
                    />
                  }
                  label="Primary key"
                />
                <FormControlLabel
                  control={
                    <Checkbox
                      checked={value.index}
                      onChange={(e) =>
                        updateForm(index, "index", e.target.checked)
                      }
                    />
                  }
                  label="Index"
                />
                <Box
                  sx={{
                    flex: 1,
                    display: "flex",
                    alignItems: "center",
                    marginLeft: 3,
                  }}
                >
                  <IconButton
                    onClick={() =>
                      setColumns(columns.filter((_, i) => i !== index))
                    }
                  >
                    <DeleteIcon fontSize="small" />
                  </IconButton>
                </Box>
              </Box>
            ))}
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
