import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

export default function Header() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar
        position="static"
        enableColorOnDark={true}
        sx={{ height: "48px", padding: "0 16px" }}
      >
        <Typography
          component="h1"
          sx={{
            flexGrow: 1,
            width: "100%",
            fontSize: "1.2rem",
            fontWeight: "bold",
            display: "flex",
            alignItems: "center",
          }}
        >
          MossAPI Admin
        </Typography>
      </AppBar>
    </Box>
  );
}
