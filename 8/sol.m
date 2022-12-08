fid = fopen('input','rt');
lines = textscan(fid, "%1n"){1};
n = sqrt(length(lines))
data = reshape(lines, n, n)';
size(data)
res = zeros(n,n);
res_tree = ones(n,n);

for k = 1:4
  for i = 1:n
    m = -1;
    for j = 1:n
      if m < data(i,j)
        m = data(i, j);
        res(i, j) = 1;
      end
      l = min(n, j+1);
      while l < n && data(i, l) < data(i, j) 
        l++;
      end
      res_tree(i,j) *= l-j;
    end
  end
  data = rot90(data);
  res = rot90(res);
  res_tree = rot90(res_tree);
end

sum(sum(res))
max(max(res_tree))

