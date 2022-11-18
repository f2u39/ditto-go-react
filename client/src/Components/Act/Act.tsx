import { useEffect, useState } from "react";
import dayjs, { Dayjs } from 'dayjs';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Link from '@mui/material/Link';
import { AppBar, DialogActions } from '@mui/material';
import { Badge } from '@mui/material';
import { Box } from '@mui/material';
import { FormControl } from '@mui/material';
import { Grid } from '@mui/material';
import { IconButton } from '@mui/material';
import { InputLabel } from '@mui/material';
import { MenuItem } from '@mui/material';
import { Select } from '@mui/material';
import { Stack } from '@mui/material';
import { Tooltip } from '@mui/material';
import { Toolbar } from '@mui/material';
import { Typography } from '@mui/material';
import PostAddIcon from '@mui/icons-material/PostAdd';
import TimerIcon from '@mui/icons-material/Timer';
import { Git as GitIcon } from 'react-bootstrap-icons';
import { Controller as ControllerIcon } from 'react-bootstrap-icons';
import SportsEsportsIcon from '@mui/icons-material/SportsEsports';
import GitHubIcon from '@mui/icons-material/GitHub';
import FormatListNumberedRtlIcon from '@mui/icons-material/FormatListNumberedRtl';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import TitleIcon from '@mui/icons-material/Title';
import ArrowCircleLeftIcon from '@mui/icons-material/ArrowCircleLeft';
import ArrowCircleRightIcon from '@mui/icons-material/ArrowCircleRight';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { DesktopDatePicker } from '@mui/x-date-pickers/DesktopDatePicker';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DatePicker, LocalizationProvider } from "@mui/x-date-pickers";

export default function Act() {
    const [date, setDate] = useState<Dayjs | null>(dayjs(new Date()))
    const [tempDate, setTempDate] = useState<Dayjs | null>(date)
    const [openNewActivity, setOpenNewActivity] = useState(false)
    const [openCalendar, setOpenCalendar] = useState(false)
    const [openStopwatch, setOpenStopwatch] = useState(false)

    const handleChangeTempDate = (newValue: Dayjs | null) => {
        setTempDate(newValue)
    }

    const handleCreateActChangeDate = (newValue: Dayjs | null) => {
        setTempDate(newValue)
        setFormCreateActValues({
            ...formCreateActValues,
            "date": newValue!!.format('YYYYMMDD'),
        })
    }

    const handleChangeTempDateSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        setDate(tempDate)
        setOpenCalendar(false)
    }

    const handlePreviousDate = () => {
        setDate(dayjs(date).add(-1, 'day'))
    }
    const handleNextDate = () => {
        setDate(dayjs(date).add(1, 'day'))
    }

    const handleNewActivityOpen = () => { setOpenNewActivity(true) }
    const handleNewActivityClose = () => {
        setFormCreateActValues(defaultCreateActValues)
        setOpenNewActivity(false)
    }

    const handleCalendarOpen = () => { setOpenCalendar(true) }
    const handleCalendarClose = () => { setOpenCalendar(false) }

    const handleStopwatchOpen = () => { setOpenStopwatch(true) }
    const handleStopwatchClose = () => {
        setFormStopwatch(defaultStopwatchValue)
        setOpenStopwatch(false)
    }

    const [acts, setActs] = useState({
        day_details: [],
        day_summary: [],
        month_details: [],
        month_summary: [],
        playing_games: [],
        stopwatch: [],
    })

    useEffect(() => {
        fetchData()
    }, [date])

    const handleStartStopwatchSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        fetch("/api/act/watch/start", {
            method: "POST",
            body: JSON.stringify(formStopwatch),
            headers: {
                "Content-Type": "application/json"
            }
        })
            .then(resp => resp.json())
            .then(data => {
                fetchData()
                handleStopwatchClose()
            })
            .catch(error => console.error("Error:", error))
    }

    const handleStopStopwatchSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        fetch("/api/act/watch/stop", {
            method: "POST",
        })
            .then(resp => resp.json())
            .then(data => {
                fetchData()
                handleStopwatchClose()
            })
    }

    function fetchData() {
        fetch(`/api/act?date=${dayjs(date).format('YYYYMMDD')}`)
            .then(resp => resp.json())
            .then(data => {
                if (data != null) {
                    setActs(data)
                }
            })
    }

    const dayDetails = Array.isArray(acts.day_details) ? acts.day_details : []
    const daySummary: any = acts.day_summary ? acts.day_summary : []
    const monDetails = Array.isArray(acts.month_details) ? acts.month_details : []
    const monSummary: any = acts.month_summary ? acts.month_summary : []
    const playingGames = Array.isArray(acts.playing_games) ? acts.playing_games : []
    const stopwatching: any = acts.stopwatch ? acts.stopwatch : []

    const defaultCreateActValues = {
        type: 'Gaming',
        date: dayjs(new Date()).format('YYYYMMDD'),
        duration: '',
        gameId: '',
    }

    const defaultStopwatchValue = {
        type: 'Gaming',
        gameId: ''
    }

    function reset() {
        setTempDate(dayjs(new Date()))
        setFormCreateActValues(defaultCreateActValues)
    }

    const [formCreateActValues, setFormCreateActValues] = useState(defaultCreateActValues)
    const [formStopwatch, setFormStopwatch] = useState(defaultStopwatchValue)

    const handleCreateActInputChange = (e: { target: { name: any; value: any; } }) => {
        const { name, value } = e.target;
        setFormCreateActValues({
            ...formCreateActValues,
            [name]: value,
        })
        console.log("CreateActInputValue: " + formCreateActValues.date)
    }

    const handleStopwatchChange = (e: { target: { name: any; value: any; } }) => {
        const { name, value } = e.target;
        setFormStopwatch({
            ...formStopwatch,
            [name]: value,
        })
    }

    const handleCreateActSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        console.log(formCreateActValues)

        fetch("/api/act/create", {
            method: "POST",
            // credentials: 'include',
            credentials: 'same-origin',
            body: JSON.stringify(formCreateActValues),
            headers: {
                "Content-Type": "application/json"
            }
        })
            .then(response => response.json())
            // .then(response => console.log("Success:", JSON.stringify(response)))
            .then(() => {
                handleNewActivityClose()
                fetchData()
                reset()
            })
            .catch(error => console.error("Error:", error))
    }

    return (
        <Box sx={{ width: '100%' }}>
            <Grid
                container
                display="flex"
                justifyContent="center"
                alignItems="center"
                xs={12}
            >
                <Grid item xs={8}>
                    <Box sx={{ flexGrow: 1 }}>
                        <AppBar
                            position="static"
                            style={{
                                background: 'transparent',
                                boxShadow: 'none'
                            }}
                        >
                            <Toolbar>
                                <Typography sx={{ flexGrow: 1 }} />
                                <IconButton
                                    size="large"
                                    aria-controls="menu-appbar"
                                    aria-haspopup="true"
                                    color="inherit"
                                >
                                    <PostAddIcon onClick={handleNewActivityOpen} sx={{ fontSize: 35 }} />
                                </IconButton>

                                <IconButton
                                    size="large"
                                    aria-controls="menu-appbar"
                                    aria-haspopup="true"
                                    color="inherit"
                                >
                                    <Badge color="secondary" badgeContent={1} invisible={acts.stopwatch === null}>
                                        <TimerIcon onClick={handleStopwatchOpen} sx={{ fontSize: 35 }} />
                                    </Badge>
                                </IconButton>

                                <Typography sx={{ flexGrow: 1 }} />
                            </Toolbar>
                        </AppBar>
                    </Box>
                    <TableContainer sx={{ border: 1, borderRadius: 1 }}>
                        <Toolbar sx={{ borderBottom: 1 }}>
                            <Tooltip title="Previous date">
                                <IconButton onClick={handlePreviousDate}>
                                    <ArrowCircleLeftIcon />
                                </IconButton>
                            </Tooltip>
                            <Grid
                                container
                                spacing={0}
                                direction="row"
                                alignItems="center"
                                justifyContent="center"
                            >
                                <Grid item>
                                    <Grid container direction="row" alignItems="center">
                                        <Link href="#" variant="body1" underline="hover" onClick={handleCalendarOpen}>
                                            {dayjs(date).format('DD MMM (ddd) YYYY')}
                                        </Link>
                                    </Grid>
                                </Grid>
                            </Grid>

                            <Tooltip title="Next date">
                                <IconButton onClick={handleNextDate}>
                                    <ArrowCircleRightIcon />
                                </IconButton>
                            </Tooltip>
                        </Toolbar>

                        <Table size="small">
                            <TableHead>
                                <TableRow>
                                    <TableCell align="center" style={{ width: 40, verticalAlign: 'top' }}><FormatListNumberedRtlIcon /></TableCell>
                                    <TableCell align="center" style={{ width: 110, verticalAlign: 'top' }}><AccessTimeIcon /></TableCell>
                                    <TableCell align="left" style={{ verticalAlign: 'top' }}><TitleIcon /></TableCell>
                                    <TableCell style={{ width: 120, verticalAlign: 'top' }}></TableCell>
                                </TableRow>
                            </TableHead>

                            <TableBody>
                                <TableRow sx={{ borderTop: 1 }}>
                                    <TableCell colSpan={4} align="center">📆Daily</TableCell>
                                </TableRow>

                                {(dayDetails).map(
                                    (detail: any) => {
                                        return (
                                            <TableRow
                                                key={detail.id}
                                                sx={{ '&:last-child td, &:last-child th': { border: 0, fontSize: 15 } }}
                                            >
                                                {detail.act.type === 'Gaming' ?
                                                    <TableCell align="center" style={{ verticalAlign: 'top' }}><Typography color="lightpink"><SportsEsportsIcon /></Typography></TableCell> :
                                                    <TableCell align="center" style={{ verticalAlign: 'top' }}><Typography color="mediumpurple"><GitHubIcon /></Typography></TableCell>}

                                                {detail.act.type === 'Gaming' ?
                                                    <TableCell align="center" style={{ verticalAlign: 'top' }}><Typography color="lightpink">{detail.hour === 0 ? '' : detail.hour + 'h'} {detail.min}m</Typography></TableCell> :
                                                    <TableCell align="center" style={{ verticalAlign: 'top' }}><Typography color="mediumpurple">{detail.hour === 0 ? '' : detail.hour + 'h'} {detail.min}m</Typography></TableCell>}

                                                {detail.act.type === 'Gaming' ?
                                                    <TableCell colSpan={2} align="left" style={{ verticalAlign: 'top' }}><Typography color="lightpink">{detail.game[0]?.title}</Typography></TableCell> :
                                                    <TableCell colSpan={2}></TableCell>}
                                            </TableRow>
                                        )
                                    }
                                )}

                                <TableRow>
                                    <TableCell colSpan={2}></TableCell>
                                    <TableCell align="right" style={{ verticalAlign: 'top' }}><Typography color="lightpink"><ControllerIcon size={23} /></Typography></TableCell>
                                    <TableCell align="right" style={{ verticalAlign: 'top' }}><Typography color="lightpink">{daySummary.game_hour === 0 ? '' : daySummary.game_hour + 'h'} {daySummary.game_min}m</Typography></TableCell>
                                </TableRow>
                                <TableRow>
                                    <TableCell colSpan={2}></TableCell>
                                    <TableCell align="right" style={{ verticalAlign: 'top' }}><Typography color="mediumpurple"><GitIcon size={23} /></Typography></TableCell>
                                    <TableCell align="right" style={{ verticalAlign: 'top' }}><Typography color="mediumpurple">{daySummary.pgm_hour === 0 ? '' : daySummary.pgm_hour + 'h'} {daySummary.pgm_min}m</Typography></TableCell>
                                </TableRow>
                                <TableRow sx={{ borderTop: 1 }}>
                                    <TableCell colSpan={4} align="center">📅Monthly</TableCell>
                                </TableRow>

                                {(monDetails).map(
                                    (detail: any) => {
                                        return (
                                            <TableRow
                                                key={detail.id}
                                                sx={{ '&:last-child td, &:last-child th': { border: 0, fontSize: 15 } }}
                                            >
                                                {detail.act.type === 'Gaming' ?
                                                    <TableCell align="center" style={{ verticalAlign: 'top' }}><Typography color="lightpink"><SportsEsportsIcon /></Typography></TableCell> :
                                                    <TableCell align="center" style={{ verticalAlign: 'top' }}><Typography color="mediumpurple"><GitHubIcon /></Typography></TableCell>}

                                                {detail.act.type === 'Gaming' ?
                                                    <TableCell align="center" style={{ verticalAlign: 'top' }}><Typography color="lightpink">{detail.hour === 0 ? '' : detail.hour + 'h'} {detail.min}m</Typography></TableCell> :
                                                    <TableCell align="center" style={{ verticalAlign: 'top' }}><Typography color="mediumpurple">{detail.hour === 0 ? '' : detail.hour + 'h'} {detail.min}m</Typography></TableCell>}

                                                {detail.act.type === 'Gaming' ?
                                                    <TableCell colSpan={2} align="left" style={{ verticalAlign: 'top' }}><Typography color="lightpink">{detail.game[0]?.title}</Typography></TableCell> :
                                                    <TableCell colSpan={2}></TableCell>}
                                            </TableRow>
                                        )
                                    }
                                )}

                                <TableRow>
                                    <TableCell colSpan={2}></TableCell>
                                    <TableCell align="right" style={{ verticalAlign: 'top' }}><Typography color="lightpink"><ControllerIcon size={23} /></Typography></TableCell>
                                    <TableCell align="right" style={{ verticalAlign: 'top' }}><Typography color="lightpink">{monSummary.game_hour === 0 ? '' : monSummary.game_hour + 'h'} {monSummary.game_min}m</Typography></TableCell>
                                </TableRow>
                                <TableRow>
                                    <TableCell colSpan={2}></TableCell>
                                    <TableCell align="right" style={{ verticalAlign: 'top' }}><Typography color="mediumpurple"><GitIcon size={23} /></Typography></TableCell>
                                    <TableCell align="right" style={{ verticalAlign: 'top' }}><Typography color="mediumpurple">{monSummary.pgm_hour === 0 ? '' : monSummary.pgm_hour + 'h'} {monSummary.pgm_min}m</Typography></TableCell>
                                </TableRow>
                            </TableBody>
                        </Table>
                    </TableContainer>
                </Grid>
            </Grid>

            <Dialog
                open={openCalendar}
                onClose={handleCalendarClose}
            >
                <DialogTitle>Select a date</DialogTitle>
                <DialogContent>
                    <form onSubmit={handleChangeTempDateSubmit}>
                        <FormControl>
                            <LocalizationProvider dateAdapter={AdapterDayjs}>
                                <DesktopDatePicker
                                    inputFormat={"MM/DD/YYYY"}
                                    value={tempDate}
                                    onChange={handleChangeTempDate}
                                    renderInput={(params) =>
                                        <TextField {...params} />
                                    }
                                />
                            </LocalizationProvider>

                            <DialogActions sx={{ mt: 1, mb: -1, mr: -1 }}>
                                <Button color="secondary" onClick={handleCalendarClose}>Cancel</Button>
                                <Button type="submit" color="success">Select</Button>
                            </DialogActions>
                        </FormControl>
                    </form>
                </DialogContent>
            </Dialog>

            <Dialog
                open={openNewActivity}
                onClose={handleNewActivityClose}
            >
                <DialogTitle align="center">New Activity</DialogTitle>
                <DialogContent>
                    <form onSubmit={handleCreateActSubmit}>
                        <FormControl
                            fullWidth
                            sx={{ mt: 2 }}
                        >
                            <InputLabel htmlFor="type">Type</InputLabel>
                            <Select
                                name="type"
                                label="Type"
                                value={formCreateActValues.type}
                                onChange={handleCreateActInputChange}
                            >
                                <MenuItem value="Gaming">Gaming</MenuItem>
                                <MenuItem value="Programming">Programming</MenuItem>
                            </Select>
                        </FormControl>

                        <FormControl
                            fullWidth
                            sx={{ mt: 2 }}
                        >
                            <LocalizationProvider dateAdapter={AdapterDayjs}>
                                <DatePicker
                                    label="Date"
                                    inputFormat={"MM/DD/YYYY"}
                                    value={tempDate}
                                    onChange={handleCreateActChangeDate}
                                    renderInput={(params) =>
                                        <TextField {...params} />
                                    }
                                />
                            </LocalizationProvider>
                        </FormControl>

                        <FormControl
                            fullWidth
                            sx={{ mt: 2 }}
                        >
                            <TextField
                                name="duration"
                                label="Duration"
                                type="number"
                                value={formCreateActValues.duration}
                                onChange={handleCreateActInputChange}
                                InputProps={{
                                    inputProps: { min: 0 }
                                }}
                            />
                        </FormControl>

                        <FormControl
                            fullWidth
                            sx={{ mt: 2 }}
                        >
                            <InputLabel htmlFor="type">Game</InputLabel>
                            <Select
                                name="gameId"
                                label="Game"
                                value={formCreateActValues.gameId}
                                inputProps={{
                                    name: 'gameId',
                                }}
                                onChange={handleCreateActInputChange}
                            >
                                {playingGames.map((game: any, index) => {
                                    return (
                                        <MenuItem key={index} value={game.id}>{game.title}</MenuItem>
                                    )
                                })}
                            </Select>
                        </FormControl>

                        <DialogActions sx={{ mt: 1, mb: -1, mr: -1 }}>
                            <Button color="secondary" onClick={handleNewActivityClose}>Cancel</Button>
                            <Button type="submit" color="success">Create</Button>
                        </DialogActions>
                    </form>
                </DialogContent>
            </Dialog>

            <Dialog
                open={openStopwatch}
                onClose={handleStopwatchClose}
            >
                <DialogTitle align="center">Stopwatch</DialogTitle>
                <DialogContent>
                    {
                        acts.stopwatch === null ?
                            <form onSubmit={handleStartStopwatchSubmit}>
                                <FormControl fullWidth sx={{ mt: 1 }}>
                                    <InputLabel htmlFor="type">Type</InputLabel>
                                    <Select
                                        name="type"
                                        label="Type"
                                        value={formStopwatch.type}
                                        onChange={handleStopwatchChange}
                                    >
                                        <MenuItem value="Gaming">Gaming</MenuItem>
                                        <MenuItem value="Programming">Programming</MenuItem>
                                    </Select>
                                </FormControl>
                                <FormControl fullWidth sx={{ mt: 2, minWidth: 250 }}>
                                    <InputLabel htmlFor="type">Game</InputLabel>
                                    <Select
                                        name="gameId"
                                        label="Game"
                                        value={formStopwatch.gameId}
                                        inputProps={{
                                            name: 'gameId',
                                        }}
                                        onChange={handleStopwatchChange}
                                    >
                                        {playingGames.map((game: any, index) => {
                                            return (
                                                <MenuItem key={index} value={game.id}>{game.title}</MenuItem>
                                            )
                                        })}
                                    </Select>
                                </FormControl>
                                <DialogActions sx={{ mt: 1, mb: -1, mr: -2 }}>
                                    <Button color="secondary" onClick={handleStopwatchClose}>Cancel</Button>
                                    <Button color="success" type="submit">Start</Button>
                                </DialogActions>
                            </form>
                        :
                            <form onSubmit={handleStopStopwatchSubmit}>
                                <FormControl fullWidth sx={{ mt: 1 }}>
                                    <TextField label="Start At" value={dayjs(stopwatching.start_time).format('YYYY/MM/DD  HH:mm:ss')} disabled></TextField>
                                </FormControl>
                                <FormControl fullWidth sx={{ mt: 2  }}>
                                    <TextField label="Type" value={stopwatching.type} disabled></TextField>
                                </FormControl>
                                <FormControl fullWidth sx={{ mt: 2, minWidth: 250 }}>
                                    <TextField label="Title" value={stopwatching.game_title} disabled></TextField>
                                </FormControl>
                                <DialogActions sx={{ mt: 1, mb: -1, mr: -1 }}>
                                    <Button color="secondary" onClick={handleStopwatchClose}>Cancel</Button>
                                    <Button type="submit" color="error">Stop</Button>
                                </DialogActions>
                            </form>
                    }
                </DialogContent>
            </Dialog>
        </Box>
    )
}
