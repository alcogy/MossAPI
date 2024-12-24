import Box from "@mui/material/Box";
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
import { Column, ColumnFormParams } from "../state/models";
import { segmentList, typeList } from "../common/constants";

interface Props {
  value: Column;
  index: number;
  updateForm: (index: number, kind: ColumnFormParams, value: any) => void;
  onClickDeleteRow: (index: number) => void;
}

export default function TableForm({
  value,
  index,
  updateForm,
  onClickDeleteRow,
}: Props) {
  return (
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
            onChange={(e) => updateForm(index, "pk", e.target.checked)}
          />
        }
        label="Primary key"
      />

      <FormControlLabel
        control={
          <Checkbox
            checked={value.nullable}
            disabled={value.pk}
            onChange={(e) => updateForm(index, "nullable", e.target.checked)}
          />
        }
        label="Null"
      />

      <FormControl sx={{ minWidth: "88px" }}>
        <InputLabel id="unique-label">Unique</InputLabel>
        <Select
          label="Unique"
          labelId="unique-label"
          defaultValue={0}
          value={value.unique}
          onChange={(e) => updateForm(index, "unique", e.target.value)}
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
  );
}
