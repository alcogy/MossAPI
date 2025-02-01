import Typography from "@mui/material/Typography";

interface Props {
  label: string;
}

export default function ModuleTitle({ label }: Props) {
  return (
    <Typography
      sx={{
        marginBottom: "16px",
        fontWeight: 700,
        fontSize: "1.2rem",
      }}
      color="textPrimary"
    >
      {label}
    </Typography>
  );
}
