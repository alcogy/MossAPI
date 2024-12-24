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
import { TableInfo } from "../state/models";
import { useParams } from "react-router-dom";
import { API_GET_TABLE_DETAIL } from "../common/constants";

export default function TableDetail() {
  const [tableInfo, setTableInfo] = useState<TableInfo>();
  const { table } = useParams();

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(API_GET_TABLE_DETAIL + table);
      const data = await response.json();
      setTableInfo(data as TableInfo);
    };
    fetchData().catch((e) => console.error(e));
  }, [table]);

  return (
    <>
      <ModuleTitle label="Table Detail" />
      <Box sx={{ marginBottom: 2 }}>
        <Typography variant="h5">{tableInfo?.tableName}</Typography>
        <Typography variant="body2" color="secondary">
          {tableInfo?.tableDesc}
        </Typography>
      </Box>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell sx={{ fontWeight: 700, width: "20%" }}>
                Column
              </TableCell>
              <TableCell sx={{ fontWeight: 700, width: "15%" }}>Type</TableCell>
              <TableCell sx={{ fontWeight: 700, width: "10%" }}>Null</TableCell>
              <TableCell sx={{ fontWeight: 700, width: "10%" }}>
                Unique
              </TableCell>
              <TableCell sx={{ fontWeight: 700, width: "10%" }}>
                Index
              </TableCell>
              <TableCell sx={{ fontWeight: 700, width: "10%" }}>PK</TableCell>
              <TableCell sx={{ fontWeight: 700, width: "25%" }}>
                Remarks
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {tableInfo?.columns?.map((v, i) => (
              <TableRow key={i}>
                <TableCell>{v.name}</TableCell>
                <TableCell>{v.type}</TableCell>
                <TableCell>{v.nullable ? "Yes" : ""}</TableCell>
                <TableCell>
                  {v.unique === 0 ? "" : `Yes(${v.unique})`}
                </TableCell>
                <TableCell>{v.index === 0 ? "" : `Yes(${v.index})`}</TableCell>
                <TableCell>{v.pk ? "Yes" : ""}</TableCell>
                <TableCell>{v.comment}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
