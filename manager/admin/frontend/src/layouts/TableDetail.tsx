import { useEffect, useState } from "react";

import {
  Box,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import Paper from "@mui/material/Paper";
import ModuleTitle from "../components/ModuleTitle";
import { Column, Table as TableModel } from "../state/models";
import { useParams } from "react-router-dom";
import { API_GET_TABLE_DETAIL } from "../common/constants";

export default function TableDetail() {
  const [tableInfo, setTableInfo] = useState<TableModel>();
  const { table } = useParams();

  const getKeyLabel = (v: Column): string => {
    if (v.pk) return "PRI";
    if (v.index) return "MUL";
    return "";
  };

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(API_GET_TABLE_DETAIL + table);
      const data = await response.json();
      setTableInfo(data as TableModel);
    };
    fetchData().catch((e) => console.error(e));
  }, [table]);

  return (
    <>
      <ModuleTitle label="Table Detail" />
      <Box sx={{ marginBottom: 1 }}>
        <Typography variant="h6">{tableInfo?.name}</Typography>
        <Typography>{tableInfo?.name}</Typography>
      </Box>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell sx={{ fontWeight: 700, width: "40%" }}>
                Column
              </TableCell>
              <TableCell sx={{ fontWeight: 700, width: "30%" }}>Type</TableCell>
              <TableCell sx={{ fontWeight: 700, width: "20%" }}>
                Not Null
              </TableCell>
              <TableCell sx={{ fontWeight: 700, width: "10%" }}>Key</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {tableInfo?.columns?.map((v, i) => (
              <TableRow key={i}>
                <TableCell>{v.name}</TableCell>
                <TableCell>{v.type}</TableCell>
                <TableCell>Yes</TableCell>
                <TableCell>{getKeyLabel(v)}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
