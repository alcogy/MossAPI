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
  Typography,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import PlayArrowIcon from "@mui/icons-material/PlayArrow";
import StopIcon from "@mui/icons-material/Stop";
import Paper from "@mui/material/Paper";
import ModuleTitle from "../components/ModuleTitle";
import AddIcon from "@mui/icons-material/Add";

const Sample = [
  {
    id: "slx7gtiurkjd",
    name: "customer",
    port: "12001",
    status: "Running",
  },
  {
    id: "ew89fjfllowh",
    name: "project",
    port: "12002",
    status: "Running",
  },
  {
    id: "ghuty47yh3uy",
    name: "relation",
    port: "12003",
    status: "Stop",
  },
];

export default function SearviceList() {
  return (
    <Paper elevation={8} sx={{ padding: "24px" }}>
      <ModuleTitle label="Service Manager" />
      <Box sx={{ marginBottom: "8px" }}>
        <Button variant="contained" startIcon={<AddIcon />}>
          Service
        </Button>
      </Box>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell sx={{ fontWeight: 700 }}>ID</TableCell>
              <TableCell sx={{ fontWeight: 700 }}>Name</TableCell>
              <TableCell sx={{ fontWeight: 700 }}>Port</TableCell>
              <TableCell sx={{ fontWeight: 700 }}>Status</TableCell>
              <TableCell sx={{ fontWeight: 700 }} align="center">
                Action
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {Sample.map((value) => (
              <TableRow key={value.id}>
                <TableCell>{value.id}</TableCell>
                <TableCell>{value.name}</TableCell>
                <TableCell>{value.port}</TableCell>
                <TableCell>{value.status}</TableCell>
                <TableCell sx={{ width: "1%", whiteSpace: "nowrap" }}>
                  <ButtonGroup
                    variant="contained"
                    aria-label="Basic button group"
                  >
                    <IconButton>
                      <PlayArrowIcon fontSize="small" />
                    </IconButton>
                    <IconButton>
                      <StopIcon fontSize="small" />
                    </IconButton>
                    <IconButton>
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
