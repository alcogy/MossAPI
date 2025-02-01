import Box from "@mui/material/Box";
import CircularProgress from "@mui/material/CircularProgress";

export default function Loading() {
  return (
    <Box
      sx={{
        width: "100vw",
        height: "100vh",
        position: "fixed",
        top: 0,
        left: 0,
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        zIndex: 99999,
        background: "rgba(0, 0, 0, 0.6)",
      }}
    >
      <CircularProgress color="secondary" />
    </Box>
  );
}
