-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和SQL文件：/Users/linxu/Projects/mygo/gopgo/server/storage/data/generate/podcast_menu.sql
-- Version: 2.16.10
-- Date: 2025-04-01 10:27:18
-- Link https://github.com/bufanyun/hotgo

-- BEGIN;

--
-- 数据库： "test_tingo_db"
--

-- --------------------------------------------------------

--
-- 插入表中的数据 "hg_admin_menu"
--

DO $$
DECLARE
    now_time timestamp := now();
    dir_id integer;
    list_id integer;
    edit_id integer;
BEGIN

-- 菜单目录
INSERT INTO "hg_admin_menu" ("id", "pid", "title", "name", "path", "icon", "type", "redirect", "permissions", "permission_name", "component", "always_show", "active_menu", "is_root", "is_frame", "frame_src", "keep_alive", "hidden", "affix", "level", "tree", "sort", "remark", "status", "created_at", "updated_at") 
VALUES (DEFAULT, '0', 'hg_podcast', 'podcast', '/podcast', 'MenuOutlined', '1', '/podcast/index', '', '', 'LAYOUT', '1', '', '0', '0', '', '0', '0', '0', '1', '', '0', '', '1', now_time, now_time)
RETURNING id INTO dir_id;

-- 菜单页面
-- 列表
INSERT INTO "hg_admin_menu" ("id", "pid", "title", "name", "path", "icon", "type", "redirect", "permissions", "permission_name", "component", "always_show", "active_menu", "is_root", "is_frame", "frame_src", "keep_alive", "hidden", "affix", "level", "tree", "sort", "remark", "status", "created_at", "updated_at") 
VALUES (DEFAULT, dir_id, 'hg_podcast列表', 'podcastIndex', 'index', '', '2', '', '/podcast/list', '', '/podcast/index', '1', 'podcast', '0', '0', '', '0', '1', '0', '2', CONCAT('tr_', dir_id,' '), '10', '', '1', now_time, now_time)
RETURNING id INTO list_id;


-- 详情
INSERT INTO "hg_admin_menu" ("id", "pid", "title", "name", "path", "icon", "type", "redirect", "permissions", "permission_name", "component", "always_show", "active_menu", "is_root", "is_frame", "frame_src", "keep_alive", "hidden", "affix", "level", "tree", "sort", "remark", "status", "created_at", "updated_at") 
VALUES (DEFAULT, list_id, 'hg_podcast详情', 'podcastView', '', '', '3', '', '/podcast/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '3', CONCAT('tr_', dir_id, ' tr_', list_id,' '), '10', '', '1', now_time, now_time);


-- 菜单按钮

-- 编辑
INSERT INTO "hg_admin_menu" ("id", "pid", "title", "name", "path", "icon", "type", "redirect", "permissions", "permission_name", "component", "always_show", "active_menu", "is_root", "is_frame", "frame_src", "keep_alive", "hidden", "affix", "level", "tree", "sort", "remark", "status", "created_at", "updated_at") 
VALUES (DEFAULT, list_id, '编辑/新增hg_podcast', 'podcastEdit', '', '', '3', '', '/podcast/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '3', CONCAT('tr_', dir_id, ' tr_', list_id,' '), '20', '', '1', now_time, now_time)
RETURNING id INTO edit_id;




-- 删除
INSERT INTO "hg_admin_menu" ("id", "pid", "title", "name", "path", "icon", "type", "redirect", "permissions", "permission_name", "component", "always_show", "active_menu", "is_root", "is_frame", "frame_src", "keep_alive", "hidden", "affix", "level", "tree", "sort", "remark", "status", "created_at", "updated_at") 
VALUES (DEFAULT, list_id, '删除hg_podcast', 'podcastDelete', '', '', '3', '', '/podcast/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '3', CONCAT('tr_', dir_id, ' tr_', list_id,' '), '40', '', '1', now_time, now_time);




-- 导出
INSERT INTO "hg_admin_menu" ("id", "pid", "title", "name", "path", "icon", "type", "redirect", "permissions", "permission_name", "component", "always_show", "active_menu", "is_root", "is_frame", "frame_src", "keep_alive", "hidden", "affix", "level", "tree", "sort", "remark", "status", "created_at", "updated_at") 
VALUES (DEFAULT, list_id, '导出hg_podcast', 'podcastExport', '', '', '3', '', '/podcast/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '3', CONCAT('tr_', dir_id, ' tr_', list_id,' '), '70', '', '1', now_time, now_time);



END $$;

COMMIT;