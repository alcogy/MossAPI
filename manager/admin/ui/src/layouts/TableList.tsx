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
import { useRecoilState } from "recoil";
import { tableListState } from "../state/atoms";
import { API_GET_TABLES, API_TABLE_DELETE } from "../common/constants";
import { useEffect } from "react";
import { TableInfo } from "../state/models";

export default function TableList() {
  const [tableList, setTableList] = useRecoilState(tableListState);

  const onClickRemove = async (table: string) => {
    const res = await fetch(API_TABLE_DELETE + table, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const json = await res.json();
    if (json["message"] === "ok") {
      await fetchTableList();
    }
  };

  const fetchTableList = async () => {
    try {
      const response = await fetch(API_GET_TABLES);
      const data = await response.json();
      setTableList(data as TableInfo[]);
    } catch (e) {
      console.error(e);
    }
  };

  useEffect(() => {
    const fetchData = async () => await fetchTableList();
    fetchData();
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
            {tableList &&
              tableList.map((value, index) => (
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
