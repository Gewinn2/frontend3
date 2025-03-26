import React from "react";
import {
  Card,
  CardContent,
  CardMedia,
  Typography,
  Button,
  Rating,
} from "@mui/material";
import { motion } from "framer-motion";
import { useDispatch } from "react-redux";
import { addToCart } from "../features/cartSlice";

const ProductCard = ({ product }) => {
  const dispatch = useDispatch();
  
  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
    >
      <Card
        sx={{
          maxWidth: 345,
          height: 400,
          margin: 2,
          boxShadow: 3,
          display: "flex",
          flexDirection: "column",
        }}
      >
        <CardMedia
          component="img"
          height="170"
          image={product.image}
          alt={product.title}
          sx={{
            width: '100%',
            objectFit: 'contain', // Сохраняет пропорции изображения
            objectPosition: 'center',
            backgroundColor: '#f5f5f5', // Светлый фон для прозрачных изображений
            borderBottom: '1px solid #e0e0e0', // Граница между изображением и контентом
            padding: '10px', // Отступы внутри картинки
            aspectRatio: '1/1', // Соотношение сторон 1:1 (квадрат)
            maxHeight: '170px', // Максимальная высота
          }}
        />
        
        <CardContent sx={{ paddingTop: 0, marginTop: "auto" }}>
          <Typography variant="h6">{product.title}</Typography>
          <Typography variant="body2">{product.description}</Typography>
          <Typography variant="h5">{product.price}рубликов</Typography>
          <Rating
            value={product.rating}
            precision={0.5}
            readOnly
            sx={{ mt: 1 }}
          />
          <Button
            variant="contained"
            onClick={() => dispatch(addToCart(product))}
            sx={{ mt: 2 }}
          >
            Добавить в корзину
          </Button>
        </CardContent>
      </Card>
    </motion.div>
  );
};

export default ProductCard;