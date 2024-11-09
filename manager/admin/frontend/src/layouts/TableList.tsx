import {
  Box,
  Button,
  ButtonGroup,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import ListAltIcon from "@mui/icons-material/ListAlt";
import Paper from "@mui/material/Paper";
import ModuleTitle from "../components/ModuleTitle";
import AddIcon from "@mui/icons-material/Add";
import { useRecoilValue } from "recoil";
import { tableListState } from "../state/atoms";
import { API_GET_TABLES, API_TABLE_DELETE } from "../common/constants";
import { useEffect, useState } from "react";

interface TableList {
  tableName: string;
  tableDesc: string;
}

export default function TableList() {
  const tableList = useRecoilValue(tableListState);
  const [tables, setTables] = useState<TableList[]>([]);

  const onClickRemove = async (table: string) => {
    const res = await fetch(API_TABLE_DELETE + table, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const json = await res.json();
    if (json["message"] === "ok") {
      setTables(tables.filter((v) => v.tableName !== table));
    }
    // TODO update table list on recoil.
  };

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(API_GET_TABLES);
      const data = await response.json();
      setTables(data as TableList[]);
    };
    fetchData().catch((e) => console.error(e));
  }, []);

  return (
    <Paper elevation={8} sx={{ padding: "24px" }}>
      <ModuleTitle label="Table Manager" />
      <Box sx={{ marginBottom: "8px" }}>
        <Button variant="contained" startIcon={<AddIcon />} href="/#/table">
          Table
        </Button>
      </Box>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell sx={{ fontWeight: 700 }}>Table</TableCell>
              <TableCell sx={{ fontWeight: 700 }}>Description</TableCell>
              <TableCell sx={{ fontWeight: 700 }} align="center">
                Action
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {tables &&
              tables.map((value, index) => (
                <TableRow key={index}>
                  <TableCell sx={{ width: 0, whiteSpace: "nowrap" }}>
                    {value.tableName}
                  </TableCell>
                  <TableCell
                    sx={{
                      whiteSpace: "nowrap",
                      overflow: "hidden",
                      textOverflow: "ellipsis",
                      maxWidth: 0,
                    }}
                  >
                    {value.tableDesc}
                  </TableCell>
                  <TableCell sx={{ width: 0, whiteSpace: "nowrap" }}>
                    <ButtonGroup
                      variant="contained"
                      aria-label="Basic button group"
                    >
                      <IconButton href={`/#/table/${value.tableName}`}>
                        <ListAltIcon fontSize="small" />
                      </IconButton>
                      <IconButton
                        onClick={() => onClickRemove(value.tableName)}
                      >
                        <DeleteIcon fontSize="small" />
                      </IconButton>
                    </ButtonGroup>
                  </TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Paper>
  );
}
