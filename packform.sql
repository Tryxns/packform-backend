PGDMP                          |            packform %   12.17 (Ubuntu 12.17-0ubuntu0.20.04.1) %   12.17 (Ubuntu 12.17-0ubuntu0.20.04.1)     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    33004    packform    DATABASE     z   CREATE DATABASE packform WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE packform;
                jevon    false                        2615    33009    packform    SCHEMA        CREATE SCHEMA packform;
    DROP SCHEMA packform;
                jevon    false            �            1259    33010    customer_companies    TABLE     m   CREATE TABLE packform.customer_companies (
    company_id integer,
    company_name character varying(16)
);
 (   DROP TABLE packform.customer_companies;
       packform         heap    postgres    false    6            �            1259    33013 	   customers    TABLE     �   CREATE TABLE packform.customers (
    user_id character varying(4),
    login character varying(4),
    password integer,
    name character varying(16),
    company_id integer,
    credit_cards character varying(32)
);
    DROP TABLE packform.customers;
       packform         heap    postgres    false    6            �            1259    33016 
   deliveries    TABLE     p   CREATE TABLE packform.deliveries (
    id integer,
    order_item_id integer,
    delivered_quantity integer
);
     DROP TABLE packform.deliveries;
       packform         heap    postgres    false    6            �            1259    33019    order_items    TABLE     �   CREATE TABLE packform.order_items (
    id integer,
    order_id integer,
    price_per_unit real,
    quantity integer,
    product character varying(16)
);
 !   DROP TABLE packform.order_items;
       packform         heap    postgres    false    6            �            1259    33022    orders    TABLE     q   CREATE TABLE packform.orders (
    id integer,
    created_at text,
    order_name text,
    customer_id text
);
    DROP TABLE packform.orders;
       packform         heap    postgres    false    6            �          0    33010    customer_companies 
   TABLE DATA           H   COPY packform.customer_companies (company_id, company_name) FROM stdin;
    packform          postgres    false    203   �       �          0    33013 	   customers 
   TABLE DATA           _   COPY packform.customers (user_id, login, password, name, company_id, credit_cards) FROM stdin;
    packform          postgres    false    204   
       �          0    33016 
   deliveries 
   TABLE DATA           M   COPY packform.deliveries (id, order_item_id, delivered_quantity) FROM stdin;
    packform          postgres    false    205   v       �          0    33019    order_items 
   TABLE DATA           X   COPY packform.order_items (id, order_id, price_per_unit, quantity, product) FROM stdin;
    packform          postgres    false    206   �       �          0    33022    orders 
   TABLE DATA           K   COPY packform.orders (id, created_at, order_name, customer_id) FROM stdin;
    packform          postgres    false    207   �       �   *   x�3��OOTPS��/�,I�2�(-����8�s��qqq ��	(      �   \   x��,K����F�&���@���/�L��4�V�]�����gjfn��U�ZR�	&LM��9�L�m��E�073����� ��j      �   P   x����0�f꘢rb�����Q� �1� ��X8��,%�Sš5Ը��Ň(��M��-O�%��&��qpΕ�G�`�D      �   �   x�e��j1E뫯���m�!U���͂Mpc�Ɔ�ϕBB�]�pvFw��1�	x�n��c������SpE/dbYD����℄�+$Z���C �RP%�(�
�#J�����(���O�훅�?�]&F5H�k�	�^����r����I��U&��h�9�����ߜt;"�D�S�0�>N�&:,	]L�+��2�a�/O�xWĻl�@������.��s���       �   �   x�3�4202�50�50
14�26�24���WP6 
zrf�%�q�U�����2��2F�eb�f�1L�	T�6M`�Lf���2��2��� ΂Ԓ".s�������z�*K����24��|��=... 6�R�     