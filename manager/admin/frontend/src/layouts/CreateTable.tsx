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

export interface ColumnForm {
  name: string;
  type: number;
  pk: boolean;
  index: boolean;
}

export default function CreateTable() {
  const initColunInfo: ColumnForm = {
    name: "",
    type: 10,
    pk: false,
    index: false,
  };
  const [tableName, setTableName] = useState<string>("");
  const [columns, setColumns] = useState<ColumnForm[]>([initColunInfo]);

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
    console.log(tableName, columns);
    return;
    await fetch("http://localhost:5500/api/table/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        table: tableName,
        columns: columns,
      }),
    });
  };

  useEffect(() => {
    return () => {
      setColumns([initColunInfo]);
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
            <IconButton onClick={() => setColumns([...columns, initColunInfo])}>
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
                    <MenuItem value={10}>int</MenuItem>
                    <MenuItem value={11}>tinyint</MenuItem>
                    <MenuItem value={20}>varchar</MenuItem>
                    <MenuItem value={23}>text</MenuItem>
                    <MenuItem value={30}>date</MenuItem>
                    <MenuItem value={31}>time</MenuItem>
                    <MenuItem value={32}>datetime</MenuItem>
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
          <Button variant="contained" onClick={onClickCreate}>
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
