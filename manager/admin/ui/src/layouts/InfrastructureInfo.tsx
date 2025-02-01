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
import CircleIcon from "@mui/icons-material/Circle";
import { useEffect, useState } from "react";
import { API_GET_INFRASTRUCTURE } from "../common/constants";

export interface Infrastructure {
  gateway: boolean;
  database: boolean;
}

export default function InfrastructureInfo() {
  const [infra, setInfra] = useState<Infrastructure>({
    gateway: false,
    database: false,
  });

  const getSignalColor = (val: boolean): "success" | "disabled" => {
    return val ? "success" : "disabled";
  };

  const getStatusLabel = (val: boolean): string => {
    return val ? "Running" : "Stop";
  };

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(API_GET_INFRASTRUCTURE);
      const data = await response.json();
      setInfra(data as Infrastructure);
    };
    fetchData().catch((e) => console.error(e));
  }, []);

  return (
    <Paper elevation={8} sx={{ padding: 3 }}>
      <ModuleTitle label="Infrastructure Info" />
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell sx={{ fontWeight: 700 }}>Module</TableCell>
              <TableCell sx={{ fontWeight: 700 }}>Status</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            <TableRow>
              <TableCell>Gateway</TableCell>
              <TableCell sx={{ width: "1%", whiteSpace: "nowrap" }}>
                <Box sx={{ display: "flex", gap: 1, alignItems: "center" }}>
                  <CircleIcon
                    color={getSignalColor(infra.gateway)}
                    fontSize="small"
                  />
                  <Typography variant="body2">
                    {getStatusLabel(infra.gateway)}
                  </Typography>
                </Box>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>MySQL</TableCell>
              <TableCell sx={{ width: "1%", whiteSpace: "nowrap" }}>
                <Box sx={{ display: "flex", gap: 1, alignItems: "center" }}>
                  <CircleIcon
                    color={getSignalColor(infra.database)}
                    fontSize="small"
                  />
                  <Typography variant="body2">
                    {getStatusLabel(infra.database)}
                  </Typography>
                </Box>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </TableContainer>
    </Paper>
  );
}
