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
import { API_TABLE_CREATE, segmentList, typeList } from "../common/constants";
import { Column, ColumnFormParams } from "../state/models";

export const initColunInfo: Column = {
  name: "",
  type: 10,
  size: 0,
  pk: false,
  notNull: true,
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
          newState["notNull"] = true;
        }
        newState["pk"] = v;
        break;
      case "notNull":
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
    const type: number = column.type;
    for (const t of typeList) {
      if (t.value === 20) return "varchar(" + column.size + ")";
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

        <TextField
          label="Table description"
          variant="outlined"
          sx={{ maxWidth: "720px" }}
          value={tableDesc}
          onChange={(e) => setTableDesc(e.target.value)}
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
                <FormControl sx={{ minWidth: "128px" }}>
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

                <TextField
                  label="Size"
                  variant="outlined"
                  value={value.size}
                  sx={{ width: "64px" }}
                  onChange={(e) => updateForm(index, "size", e.target.value)}
                  disabled={![20].includes(value.type)}
                />

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
                      checked={value.notNull}
                      disabled={value.pk}
                      onChange={(e) =>
                        updateForm(index, "notNull", e.target.checked)
                      }
                    />
                  }
                  label="Not Null"
                />

                <FormControl sx={{ minWidth: "88px" }}>
                  <InputLabel id="unique-label">Unique</InputLabel>
                  <Select
                    label="Unique"
                    labelId="unique-label"
                    defaultValue={0}
                    value={value.unique}
                    onChange={(e) =>
                      updateForm(index, "unique", e.target.value)
                    }
                  >
                    {segmentList.map((v) => (
                      <MenuItem value={v.value} key={v.value}>
                        {v.label}
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>

                <FormControl sx={{ minWidth: "88px" }}>
                  <InputLabel id="index-label">Index</InputLabel>
                  <Select
                    label="Index"
                    labelId="index-label"
                    defaultValue={0}
                    value={value.index}
                    onChange={(e) => updateForm(index, "index", e.target.value)}
                  >
                    {segmentList.map((v) => (
                      <MenuItem value={v.value} key={v.value}>
                        {v.label}
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>

                <TextField
                  label="Remarks"
                  variant="outlined"
                  value={value.comment}
                  sx={{ width: "256px" }}
                  onChange={(e) => updateForm(index, "comment", e.target.value)}
                />

                <Box
                  sx={{
                    flex: 1,
                    display: "flex",
                    alignItems: "center",
                    marginLeft: 3,
                  }}
                >
                  <IconButton onClick={() => onClickDeleteRow(index)}>
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
