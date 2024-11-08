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

// TODO only ui.
export default function InfrastructureInfo() {
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
                  <CircleIcon color="success" fontSize="small" />
                  <Typography variant="body2">Running</Typography>
                </Box>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>MySQL</TableCell>
              <TableCell sx={{ width: "1%", whiteSpace: "nowrap" }}>
                <Box sx={{ display: "flex", gap: 1, alignItems: "center" }}>
                  <CircleIcon color="error" fontSize="small" />
                  <Typography variant="body2">Stop</Typography>
                </Box>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </TableContainer>
    </Paper>
  );
}
