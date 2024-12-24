import Box from "@mui/material/Box";
import ModuleTitle from "../components/ModuleTitle";
import ViewModuleIcon from "@mui/icons-material/ViewModule";
import { Button, Typography } from "@mui/material";
import StorageIcon from "@mui/icons-material/Storage";

export default function BlankContent() {
  return (
    <Box>
      <ModuleTitle label="Select content." />
      <Box sx={{ display: "flex", gap: 2 }}>
        <Button color="inherit" href="/#/service">
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
            }}
          >
            <ViewModuleIcon fontSize="large" />
            <Typography variant="caption">Add service</Typography>
          </Box>
        </Button>

        <Button color="inherit" href="/#/table">
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
            }}
          >
            <StorageIcon fontSize="large" />
            <Typography variant="caption">Add table</Typography>
          </Box>
        </Button>
      </Box>
    </Box>
  );
}
