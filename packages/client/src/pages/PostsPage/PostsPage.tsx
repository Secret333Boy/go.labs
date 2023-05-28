import React, { useEffect, useState } from "react";
import ApiService from "../../services/ApiService";
import {
  Button,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { Post } from "../../types";

const PostsPage = () => {
  const navigate = useNavigate();

  const [posts, setPosts] = useState<Post[]>([]);

  const loadPosts = async () => {
    const { data } = await ApiService.get<Post[]>("/posts");

    if (data) setPosts(data);
  };

  useEffect(() => {
    loadPosts();
  }, []);

  return (
    <div className="w-full h-screen min-h-screen">
      <div className="flex justify-end">
        <Button onClick={() => navigate("/posts/create")}>Create post</Button>
      </div>

      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }}>
          <TableHead>
            <TableRow>
              <TableCell align="center">Title</TableCell>
              <TableCell align="center">Description</TableCell>
              <TableCell align="center">Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {posts.map((post) => (
              <TableRow key={post.Id}>
                <TableCell component="th" scope="row" align="center">
                  {post.Title}
                </TableCell>
                <TableCell align="center">{post.Description}</TableCell>
                <TableCell align="center">sadfsdf</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
};

export default PostsPage;
