import * as React from 'react';
import { styled } from '@mui/material/styles';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import Collapse from '@mui/material/Collapse';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import TuneIcon from '@mui/icons-material/Tune';
import PlayCircleOutlineIcon from '@mui/icons-material/PlayCircleOutline';
import { Badge, Box, Grid, InputAdornment, Tabs, TextField } from '@mui/material';
import Pagination from '@mui/material/Pagination';

import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';

import { Check2Square, Tablet, PcDisplay, NintendoSwitch, Playstation, Xbox, BatteryHalf } from 'react-bootstrap-icons';
import { Code, CodeSlash } from 'react-bootstrap-icons';
import { Battery, BatteryCharging, BatteryFull } from 'react-bootstrap-icons';
import { useEffect, useState } from 'react';

interface ExpandMoreProps extends IconButtonProps {
    expand: boolean;
}

const ExpandMore = styled((props: ExpandMoreProps) => {
    const { expand, ...other } = props;
    return <IconButton {...other} />;
})(({ theme, expand }) => ({
    transform: !expand ? 'rotate(0deg)' : 'rotate(180deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
        duration: theme.transitions.duration.shortest,
    }),
}));

export default function Game() {
    const [details, setDetails] = useState<Detail[]>([]);
    const [platform, setPlatform] = useState('PC');
    const [page, setPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);
    const [expandedId, setExpandedId] = useState(-1)
    const [status, setStatus] = useState("Playing")
    const [playedCount, setPlayedCount] = useState(0);
    const [playingCount, setPlayingCount] = useState(0);
    const [blockingCount, setBlockingCount] = useState(0);

    useEffect(() => {
        fetch(`/api/game/status/${status}/${platform}/${page}`)
            .then(response => response.json())
            .then(data => {
                console.log(data["details"])
                if (data["details"] != null) {
                    setDetails(data["details"])
                } else {
                    setDetails([])
                }

                setTotalPages(data["total_pages"])
            })

        fetch(`/api/game/counts`)
            .then(resp => resp.json())
            .then(data => {
                setPlayedCount(data["played_cnt"])
                setPlayingCount(data["playing_cnt"])
                setBlockingCount(data["blocking_cnt"])
            })
    }, [status, platform, page]);

    const handleExpandClick = (i: number) => {
        setExpandedId(expandedId === i ? -1 : i);
    }

    const handlePageChange = (event: React.ChangeEvent<unknown>, value: number) => {
        setExpandedId(-1);
        setPage(value);
    };

    const handleStatusChange = (event: React.SyntheticEvent, newStatus: string) => {
        setExpandedId(-1);
        setPage(1);
        setStatus(newStatus);
    };

    const handlePlatformChange = (event: React.SyntheticEvent, newValue: string) => {
        setExpandedId(-1);
        setPage(1);
        setPlatform(newValue);
    };

    return (
        <Box sx={{ width: '100%' }}>
            <TabContext value={status}>
                <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                    <TabList indicatorColor="secondary" onChange={handleStatusChange} centered>
                        <Tab
                            icon={
                                <Badge badgeContent={playedCount} color="primary">
                                    <BatteryFull fontSize="30" color="white" />
                                </Badge>
                            }
                            value="Played"
                        />

                        <Tab
                            icon={
                                <Badge badgeContent={playingCount} color="success">
                                    <BatteryCharging fontSize="30" color="green" />
                                </Badge>
                            }
                            value="Playing"
                        />

                        <Tab
                            icon={
                                <Badge badgeContent={blockingCount} color="error">
                                    <Battery fontSize="30" color="red" />
                                </Badge>
                            }
                            value="Blocking"
                        />
                    </TabList>
                </Box>

                <TabPanel value={status}>
                    <Grid
                        container
                        direction="row"
                        justifyContent="space-between"
                        sx={{
                            display: 'inline-flex',
                        }}
                    >
                        <Grid
                            item
                            sx={{borderRight: 1, borderColor: 'divider' }}
                        >
                            <Tabs
                                variant="fullWidth"
                                orientation="vertical"
                                value={platform}
                                onChange={handlePlatformChange}
                            >
                                <Tab icon={<Check2Square color="white" size={30} />} value="All" />
                                <Tab icon={<PcDisplay color="orange" size={30} />} value="PC" />
                                <Tab icon={<Playstation color="skyblue" size={30} />} value="PlayStation" />
                                <Tab icon={<NintendoSwitch color="red" size={30} />} value="Nintendo Switch" />
                                <Tab icon={<Xbox color="green" size={30} />} value="Xbox" />
                                <Tab icon={<Tablet color="purple" size={30} />} value="Mobile" />
                            </Tabs>
                        </Grid>
                        <Grid item xs={10}>
                            <Grid
                                container
                            // justifyContent="center"
                            // sx={{ flexGrow: 1 }}
                            >
                                {details.map((element, i) => (
                                    <Card
                                        sx={{ ml: 3, mt: 3, maxWidth: 300 }}
                                        key={element.game.id}
                                    >
                                        <CardMedia
                                            component="img"
                                            height="300"
                                            image={"static/images/games/" + element.game.id + ".webp"}
                                        />
                                        <CardContent>
                                            <Typography variant="body2" color="text.secondary">
                                                {element.game.title}
                                            </Typography>
                                        </CardContent>
                                        <CardActions disableSpacing>
                                            <IconButton>
                                                <TuneIcon />
                                            </IconButton>

                                            <IconButton>
                                                <PlayCircleOutlineIcon />
                                            </IconButton>
                                            <ExpandMore
                                                expand={expandedId === i}
                                                onClick={() => handleExpandClick(i)}
                                            >
                                                <ExpandMoreIcon />
                                            </ExpandMore>
                                        </CardActions>
                                        <Collapse in={expandedId === i} timeout="auto" unmountOnExit>
                                            <CardContent>
                                                <Box sx={{
                                                    mx: "auto",
                                                    display: 'flex',
                                                    flexDirection: 'column',
                                                    alignItems: 'center',
                                                    '& > *': {
                                                        m: 1,
                                                    },
                                                }}
                                                >
                                                    <TextField
                                                        disabled
                                                        inputProps={{
                                                            style: { padding: '7px 5px', textAlign: 'right' },
                                                        }}
                                                        InputProps={{
                                                            startAdornment: (
                                                                <InputAdornment position="start">
                                                                    {element.game.platform === 'Mobile' ? <Tablet /> : <></>}
                                                                    {element.game.platform === 'PC' ? <PcDisplay /> : <></>}
                                                                    {element.game.platform === 'Playstation' ? <Playstation /> : <></>}
                                                                    {element.game.platform === 'Nintendo Switch' ? <NintendoSwitch /> : <></>}
                                                                    {element.game.platform === 'Xbox' ? <Xbox /> : <></>}
                                                                </InputAdornment>
                                                            ),
                                                            endAdornment: (
                                                                <InputAdornment position="end">
                                                                    {element.game.rating}
                                                                </InputAdornment>
                                                            )
                                                        }}
                                                    />
                                                    <TextField
                                                        disabled
                                                        sx={{ pt: 1 }}
                                                        inputProps={{
                                                            style: { padding: '7px 5px', textAlign: 'right' },
                                                        }}
                                                        InputProps={{
                                                            startAdornment: (
                                                                <InputAdornment position="start">
                                                                    <Code />
                                                                </InputAdornment>
                                                            ),
                                                            endAdornment: (
                                                                <InputAdornment position="end">
                                                                    {element.developer.name}
                                                                </InputAdornment>
                                                            )
                                                        }}
                                                    />

                                                    <TextField
                                                        disabled
                                                        sx={{ pt: 1 }}
                                                        inputProps={{
                                                            style: { padding: '7px 5px', textAlign: 'right' },
                                                        }}
                                                        InputProps={{
                                                            startAdornment: (
                                                                <InputAdornment position="start">
                                                                    <CodeSlash />
                                                                </InputAdornment>
                                                            ),
                                                            endAdornment: (
                                                                <InputAdornment position="end">
                                                                    {element.publisher.name}
                                                                </InputAdornment>
                                                            )
                                                        }}
                                                    />

                                                    <TextField
                                                        disabled
                                                        sx={{ pt: 1 }}
                                                        inputProps={{
                                                            style: { padding: '7px 5px', textAlign: 'right' },
                                                        }}
                                                        InputProps={{
                                                            startAdornment: (
                                                                <InputAdornment position="start">
                                                                    {element.game.status === 'Played' ? <BatteryFull /> : <></>}
                                                                    {element.game.status === 'Playing' ? <BatteryCharging /> : <></>}
                                                                    {element.game.status === 'Blocking' ? <Battery /> : <></>}
                                                                </InputAdornment>
                                                            ),
                                                            endAdornment: (
                                                                <InputAdornment position="end">
                                                                    {element.play_hour}h
                                                                </InputAdornment>
                                                            )
                                                        }}
                                                    />
                                                </Box>
                                            </CardContent>
                                        </Collapse>
                                    </Card>
                                ))
                                }
                            </Grid>
                        </Grid>
                        <Grid xs={12} sx={{ pt: 3, pb: 3 }}>
                            <Box
                                display="flex"
                                justifyContent="center"
                                alignItems="center"
                            >
                                <Pagination
                                    count={totalPages}
                                    page={page}
                                    onChange={handlePageChange}
                                    variant="outlined"
                                    color="secondary" />
                            </Box>
                        </Grid>
                    </Grid>
                </TabPanel>
            </TabContext>
        </Box>
    );

    async function fetchDetails() {
        try {
            let url = `/api/game/status/Played/${page}`
            let resp = await fetch(url);
            let json = await resp.json();
            setDetails(json["details"]);
            setTotalPages(json["total_pages"])
        } catch (error) {
            console.log(error);
        }
    }
}

function a11yProps(index: number) {
    return {
        id: `vertical-tab-${index}`,
        'aria-controls': `vertical-tabpanel-${index}`,
    };
}

function GameTabPanel() {
    return (
        <div
            role="tabpanel"
        >
            <Box sx={{ p: 3 }}>

            </Box>
        </div>
    );
}

interface Detail {
    game: Game,
    developer: Developer,
    publisher: Publisher,
    play_hour: 0,
    play_min: 0,
}

interface Game {
    id: string,
    title: string,
    genre: string,
    platform: string,
    developer_id: string,
    publisher_id: string,
    status: string,
    playtime: string,
    rating: string,
}

interface Developer {
    id: string,
    name: string,
}

interface Publisher {
    id: string,
    name: string,
}