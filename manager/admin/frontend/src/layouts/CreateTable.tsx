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
import AddIcon from "@mui/icons-material/Add";
import { useState } from "react";

export interface ColumnInfo {
  name: string;
  type: number;
  pk: boolean;
  index: boolean;
}

export const initColunInfo = { name: "", type: 0, pk: false, index: false };

export default function CreateTable() {
  const [columns, setColumns] = useState<ColumnInfo[]>([initColunInfo]);
  const onClickCreate = async () => {
    // const response = await fetch("http://localhost:5500/api/containers");
    // const json = await response.json();
    // //  const body = await reader.text();
    // console.log(json);
    await fetch("http://localhost:5500/api/container", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        service: "customer",
        port: "12090",
        artifact:
          "C:\\Users\\info\\Dev\\modular-synthesis-api\\samples\\app\\output",
      }),
    });
  };
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
        />

        <Box>
          <Box
            sx={{
              display: "flex",
              gap: 10,
              alignItems: "center",
              justifyContent: "space-between",
              maxWidth: "640px",
              marginBottom: 3,
            }}
          >
            <Typography variant="subtitle1" sx={{ fontWeight: "bold" }}>
              Setting Columns
            </Typography>
            <Button
              variant="contained"
              startIcon={<AddIcon />}
              onClick={() => setColumns([...columns, initColunInfo])}
            >
              Column
            </Button>
          </Box>
          <Box sx={{ display: "flex", flexDirection: "column", gap: 3 }}>
            {columns.map((value, index) => (
              <Box sx={{ display: "flex", gap: 1 }} key={index}>
                <TextField label="Column name" variant="outlined" />
                <FormControl sx={{ minWidth: "128px" }}>
                  <InputLabel id="select-label">Column Type</InputLabel>
                  <Select
                    label="Column Type"
                    labelId="select-label"
                    defaultValue={10}
                    onChange={(e) => console.log(e)}
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
                <FormControlLabel control={<Checkbox />} label="Primary key" />
                <FormControlLabel control={<Checkbox />} label="Index" />
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
          <Button variant="contained">Create</Button>
          <Button variant="outlined" color="secondary">
            Cancel
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
